CREATE TABLE IF NOT EXISTS lessons (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    board_id BIGINT,
    meet_link VARCHAR(255),
    start_time TIMESTAMPTZ DEFAULT now(),
    duration BIGINT,
    tutor_id BIGINT NOT NULL,
    status status_enum DEFAULT 'SCHEDULED' NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);