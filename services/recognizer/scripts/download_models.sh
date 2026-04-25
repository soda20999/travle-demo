#!/bin/bash
set -e

MODEL_DIR="${1:-./models}"
mkdir -p "$MODEL_DIR"

echo "Downloading PicoDet mainbody detection model..."
wget -q -O "$MODEL_DIR/picodet.tar" \
  "https://paddle-imagenet-models-name.bj.bcebos.com/dygraph/rec/models/inference/picodet_PPLCNet_x2_5_mainbody_lite_v1.0_infer.tar"
tar -xf "$MODEL_DIR/picodet.tar" -C "$MODEL_DIR"
mv "$MODEL_DIR/picodet_PPLCNet_x2_5_mainbody_lite_v1.0_infer" \
   "$MODEL_DIR/picodet_PPLCNet_x2_5_mainbody_lite_v1.0"
rm "$MODEL_DIR/picodet.tar"

echo "Downloading PP-ShiTuV2 feature extraction model..."
wget -q -O "$MODEL_DIR/rec.tar" \
  "https://paddle-imagenet-models-name.bj.bcebos.com/dygraph/rec/models/inference/PP-ShiTuV2/general_PPLCNetV2_base_pretrained_v1.0_infer.tar"
tar -xf "$MODEL_DIR/rec.tar" -C "$MODEL_DIR"
mv "$MODEL_DIR/general_PPLCNetV2_base_pretrained_v1.0_infer" \
   "$MODEL_DIR/general_PPLCNetV2_base_pretrained_v1.0"
rm "$MODEL_DIR/rec.tar"

echo "Models downloaded to $MODEL_DIR"
ls -la "$MODEL_DIR"/*/
