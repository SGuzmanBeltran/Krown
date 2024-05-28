-- +goose Up
CREATE TABLE tournaments (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    entry_fee BIGINT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    is_finished BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE tournaments;