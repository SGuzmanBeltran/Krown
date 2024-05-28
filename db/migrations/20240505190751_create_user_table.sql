-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    teams_count INT DEFAULT 0
);

-- +goose Down
DROP TABLE users;