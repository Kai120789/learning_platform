CREATE TABLE IF NOT EXISTS lesson_medias (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    lesson_id BIGINT REFERENCES lessons(id) NOT NULL,
    s3_link VARCHAR(255) NOT NULL,
    s3_preview VARCHAR(255) NOT NULL,
    type type_enum NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);