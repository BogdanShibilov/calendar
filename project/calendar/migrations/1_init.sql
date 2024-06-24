-- +goose Up
CREATE TABLE IF NOT EXISTS events
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    start_time TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS events;
