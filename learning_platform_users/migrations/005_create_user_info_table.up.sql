CREATE TABLE IF NOT EXISTS user_info (
    user_id BIGINT PRIMARY KEY REFERENCES users(id),
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50),
    city VARCHAR(50),
    about TEXT,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);
