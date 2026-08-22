CREATE TABLE IF NOT EXISTS tutor_subjects (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tutor_id BIGINT NOT NULL,
    subject_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE (tutor_id, subject_id)
);