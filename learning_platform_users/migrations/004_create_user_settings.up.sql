CREATE TABLE IF NOT EXISTS user_settings (
    user_id BIGINT PRIMARY KEY REFERENCES users(id),
    is_2fa_enabled BOOLEAN DEFAULT false NOT NULL,
    is_notifications_enabled BOOLEAN DEFAULT false  NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);