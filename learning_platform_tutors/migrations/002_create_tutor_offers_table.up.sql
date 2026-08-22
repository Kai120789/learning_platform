CREATE TABLE IF NOT EXISTS tutor_offers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tutor_id BIGINT NOT NULL,
    subject_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    price BIGINT NOT NULL,
    duration_minutes BIGINT,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);