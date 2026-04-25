import os

import psycopg
from psycopg.rows import dict_row

DATABASE_URL = os.environ.get(
    "DATABASE_URL",
    "postgresql://travle:travle@localhost:5432/travle",
)


def get_connection():
    return psycopg.connect(DATABASE_URL, row_factory=dict_row)


def load_all_features() -> list[dict]:
    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(
                "SELECT id, image_url, feature_vector "
                "FROM attraction_images "
                "WHERE feature_vector IS NOT NULL"
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


def get_image_url(image_id: int) -> str | None:
    with get_connection() as conn:
        with conn.cursor() as cur:
            cur.execute(
                "SELECT image_url FROM attraction_images WHERE id = %s",
                (image_id,),
            )
            row = cur.fetchone()
            return row["image_url"] if row else None
