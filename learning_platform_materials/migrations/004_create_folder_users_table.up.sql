CREATE TABLE IF NOT EXISTS folder_users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL,
    folder_id BIGINT NOT NULL REFERENCES material_folders(id)ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    UNIQUE (folder_id, user_id)
);