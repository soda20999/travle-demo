import threading

import faiss
import numpy as np

from app import db, predict

FEATURE_DIM = 512

_lock = threading.Lock()
_index: faiss.IndexIDMap | None = None
_id_list: list[int] = []


def _build_index(ids: list[int], vectors: list[np.ndarray]) -> faiss.IndexIDMap:
    base_index = faiss.IndexFlatL2(FEATURE_DIM)
    index = faiss.IndexIDMap(base_index)
    if ids:
        id_array = np.array(ids, dtype=np.int64)
        vec_array = np.vstack(vectors).astype(np.float32)
        index.add_with_ids(vec_array, id_array)
    return index


def load_index():
    global _index, _id_list
    rows = db.load_all_images()
    ids = []
    vectors = []
    for row in rows:
        if row["feature_vector"] is not None:
            vec = np.frombuffer(row["feature_vector"], dtype=np.float32)
        else:
            vec = predict.extract_feature_from_path(row["image_path"])
            db.save_feature_vector(row["id"], vec.tobytes())
        ids.append(row["id"])
        vectors.append(vec)

    with _lock:
        _index = _build_index(ids, vectors)
        _id_list = ids


def search(query_vector: np.ndarray, top_k: int = 5) -> list[dict]:
    with _lock:
        if _index is None or _index.ntotal == 0:
            return []
        query = query_vector.reshape(1, -1).astype(np.float32)
        k = min(top_k, _index.ntotal)
        distances, indices = _index.search(query, k)

    results = []
    for i in range(k):
        idx = int(indices[0][i])
        if idx == -1:
            continue
        dist = float(distances[0][i])
        similarity = 1.0 / (1.0 + dist)
        results.append({"image_id": idx, "similarity": round(similarity, 4)})
    return results


def add_to_index(image_id: int, feature_vector: np.ndarray):
    with _lock:
        global _index, _id_list
        if _index is None:
            _index = _build_index([], [])
            _id_list = []
        id_array = np.array([image_id], dtype=np.int64)
        vec_array = feature_vector.reshape(1, -1).astype(np.float32)
        _index.add_with_ids(vec_array, id_array)
        _id_list.append(image_id)


def remove_from_index(image_id: int):
    global _index, _id_list
    with _lock:
        if _index is None:
            return
        if image_id not in _id_list:
            return
        new_id_list = [i for i in _id_list if i != image_id]

    rows = db.load_all_images()
    id_set = set(new_id_list)
    ids = []
    vectors = []
    for row in rows:
        if row["id"] in id_set:
            ids.append(row["id"])
            vec = np.frombuffer(row["feature_vector"], dtype=np.float32)
            vectors.append(vec)

    new_index = _build_index(ids, vectors)
    with _lock:
        _index = new_index
        _id_list = new_id_list


def rebuild_index():
    load_index()
