CREATE TABLE IF NOT EXISTS user_info (
    user_id BIGINT PRIMARY KEY REFERENCES users(id),
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    tg_username VARCHAR(250),
    city VARCHAR(100),
    about TEXT,
    avatar VARCHAR(250),
    gender gender_enum DEFAULT 'UNKNOWN',
    birth_date DATE DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);
