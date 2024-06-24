-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;