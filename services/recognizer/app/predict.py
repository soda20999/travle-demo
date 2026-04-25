import os
from pathlib import Path

import numpy as np
from PIL import Image
from paddle import inference as paddle_infer

MODEL_DIR = Path(os.environ.get("MODEL_DIR", "/app/models"))
DET_MODEL_DIR = MODEL_DIR / "picodet_PPLCNet_x2_5_mainbody_lite_v1.0"
REC_MODEL_DIR = MODEL_DIR / "general_PPLCNetV2_base_pretrained_v1.0"

_det_predictor = None
_rec_predictor = None


def _create_predictor(model_dir: Path) -> paddle_infer.Predictor:
    model_file = str(model_dir / "inference.pdmodel")
    params_file = str(model_dir / "inference.pdiparams")
    config = paddle_infer.Config(model_file, params_file)
    config.disable_gpu()
    config.enable_mkldnn()
    config.set_cpu_math_library_num_threads(4)
    config.switch_ir_optim(True)
    config.enable_memory_optim()
    return paddle_infer.create_predictor(config)


def _get_det_predictor() -> paddle_infer.Predictor:
    global _det_predictor
    if _det_predictor is None:
        _det_predictor = _create_predictor(DET_MODEL_DIR)
    return _det_predictor


def _get_rec_predictor() -> paddle_infer.Predictor:
    global _rec_predictor
    if _rec_predictor is None:
        _rec_predictor = _create_predictor(REC_MODEL_DIR)
    return _rec_predictor


def _preprocess_det(img: Image.Image) -> np.ndarray:
    target_size = 640
    img_resized = img.resize((target_size, target_size))
    data = np.array(img_resized, dtype=np.float32)
    data = data / 255.0
    mean = np.array([0.485, 0.456, 0.406], dtype=np.float32)
    std = np.array([0.229, 0.224, 0.225], dtype=np.float32)
    data = (data - mean) / std
    data = data.transpose((2, 0, 1))
    return data[np.newaxis, :]


def _detect_mainbody(img: Image.Image) -> tuple[int, int, int, int]:
    predictor = _get_det_predictor()
    input_data = _preprocess_det(img)

    input_names = predictor.get_input_names()
    input_handle = predictor.get_input_handle(input_names[0])
    input_handle.reshape(input_data.shape)
    input_handle.copy_from_cpu(input_data)

    if len(input_names) > 1:
        scale_handle = predictor.get_input_handle(input_names[1])
        scale_factor = np.array([[img.height / 640.0, img.width / 640.0]], dtype=np.float32)
        scale_handle.reshape(scale_factor.shape)
        scale_handle.copy_from_cpu(scale_factor)

    predictor.run()

    output_names = predictor.get_output_names()
    boxes_handle = predictor.get_output_handle(output_names[0])
    boxes = boxes_handle.copy_to_cpu()

    if boxes.shape[0] == 0 or boxes.ndim < 2:
        return 0, 0, img.width, img.height

    best = boxes[0]
    x1 = max(0, int(best[2]))
    y1 = max(0, int(best[3]))
    x2 = min(img.width, int(best[4]))
    y2 = min(img.height, int(best[5]))

    if (x2 - x1) < 10 or (y2 - y1) < 10:
        return 0, 0, img.width, img.height
    return x1, y1, x2, y2


def _preprocess_rec(img: Image.Image) -> np.ndarray:
    target_size = (224, 224)
    img_resized = img.resize(target_size)
    data = np.array(img_resized, dtype=np.float32)
    data = data / 255.0
    mean = np.array([0.485, 0.456, 0.406], dtype=np.float32)
    std = np.array([0.229, 0.224, 0.225], dtype=np.float32)
    data = (data - mean) / std
    data = data.transpose((2, 0, 1))
    return data[np.newaxis, :]


def extract_feature(img: Image.Image) -> np.ndarray:
    x1, y1, x2, y2 = _detect_mainbody(img)
    cropped = img.crop((x1, y1, x2, y2))

    predictor = _get_rec_predictor()
    input_data = _preprocess_rec(cropped)

    input_names = predictor.get_input_names()
    input_handle = predictor.get_input_handle(input_names[0])
    input_handle.reshape(input_data.shape)
    input_handle.copy_from_cpu(input_data)

    predictor.run()

    output_names = predictor.get_output_names()
    output_handle = predictor.get_output_handle(output_names[0])
    feature = output_handle.copy_to_cpu()

    norm = np.linalg.norm(feature)
    if norm > 0:
        feature = feature / norm
    return feature.flatten().astype(np.float32)


def extract_feature_from_path(image_path: str) -> np.ndarray:
    img = Image.open(image_path).convert("RGB")
    return extract_feature(img)
