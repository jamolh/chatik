
CREATE TABLE users
(
    id UUID PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    username VARCHAR(64) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(), -- TIMESTAMP in UTC
    updated_at TIMESTAMPTZ
)