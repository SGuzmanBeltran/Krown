-- +goose Up
CREATE TABLE tournaments (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    entry_fee INT NOT NULL,
    start_time TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE tournaments;