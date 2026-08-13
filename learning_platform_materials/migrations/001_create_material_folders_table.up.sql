CREATE TABLE IF NOT EXISTS material_folders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    parent_folder_id BIGINT REFERENCES material_folders(id) ON DELETE CASCADE,
    tutor_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE (parent_folder_id, title, tutor_id)
);