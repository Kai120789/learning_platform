CREATE TABLE IF NOT EXISTS schedule_slots (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    schedule_id BIGINT REFERENCES schedules(id) NOT NULL,
    start_time TIMESTAMPTZ DEFAULT now(),
    status status_enum DEFAULT 'FREE' NOT NULL,
    duration BIGINT,
    lesson_id BIGINT,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);