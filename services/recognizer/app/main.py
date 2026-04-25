import io
from contextlib import asynccontextmanager

from fastapi import FastAPI, File, Form, UploadFile
from pydantic import BaseModel
from PIL import Image

from app import predict, index_manager, db


@asynccontextmanager
async def lifespan(app: FastAPI):
    index_manager.load_index()
    yield


app = FastAPI(lifespan=lifespan)


@app.get("/health")
def health():
    return {"status": "ok"}


@app.post("/predict")
def predict_endpoint(
    image: UploadFile = File(...),
    top_k: int = Form(default=5),
):
    contents = image.file.read()
    img = Image.open(io.BytesIO(contents)).convert("RGB")
    feature = predict.extract_feature(img)
    results = index_manager.search(feature, top_k=top_k)
    return {"results": results}


class IndexAddRequest(BaseModel):
    image_id: int
    image_url: str


@app.post("/index/add")
def index_add(req: IndexAddRequest):
    feature = predict.extract_feature_from_url(req.image_url)
    feature_bytes = feature.tobytes()
    db.save_feature_vector(req.image_id, feature_bytes)
    index_manager.add_to_index(req.image_id, feature)
    return {"status": "ok"}


class IndexRemoveRequest(BaseModel):
    image_id: int


@app.post("/index/remove")
def index_remove(req: IndexRemoveRequest):
    index_manager.remove_from_index(req.image_id)
    return {"status": "ok"}


@app.post("/index/rebuild")
def index_rebuild():
    index_manager.rebuild_index()
    return {"status": "ok"}
