CREATE TABLE IF NOT EXISTS tutor_reviews (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tutor_id BIGINT NOT NULL,
    author_id BIGINT NOT NULL,
    subject_id BIGINT NOT NULL,
    text TEXT NOT NULL,
    rating SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE (author_id, tutor_id)
);