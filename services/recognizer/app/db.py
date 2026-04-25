import os

import psycopg
from psycopg.rows import dict_row

DATABASE_URL = os.environ.get(
    "DATABASE_URL",
    "postgresql://travle:travle@localhost:5432/travle",
)


def get_connection():
    return psycopg.connect(DATABASE_URL, row_factory=dict_row)


def load_all_images() -> list[dict]:
    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(
                "SELECT id, image_path, feature_vector "
                "FROM attraction_images"
            )
            return cur.fetchall()


def save_feature_vector(image_id: int, feature_vector: bytes):
    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(
                "UPDATE attraction_images SET feature_vector = %s WHERE id = %s",
                (feature_vector, image_id),
            )
        conn.commit()


def get_image_path(image_id: int) -> str | None:
    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(
                "SELECT image_path FROM attraction_images WHERE id = %s",
                (image_id,),
            )
            row = cur.fetchone()
            return row["image_path"] if row else None
