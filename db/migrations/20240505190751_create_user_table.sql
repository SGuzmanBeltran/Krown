-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username text NOT NULL,
    email text NOT NULL,
    password text NOT NULL
);

-- +goose Down
DROP TABLE users;