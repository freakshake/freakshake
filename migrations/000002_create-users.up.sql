CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    avatar TEXT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    mobile_number TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT(NOW()),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
