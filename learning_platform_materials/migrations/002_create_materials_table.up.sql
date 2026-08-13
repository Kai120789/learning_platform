CREATE TABLE IF NOT EXISTS materials (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    size BIGINT NOT NULL,
    tutor_id BIGINT NOT NULL,
    folder_id BIGINT REFERENCES material_folders(id) ON DELETE CASCADE,
    mime_type VARCHAR(250) NOT NULL,
    media_object_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE (folder_id, title, tutor_id)
);