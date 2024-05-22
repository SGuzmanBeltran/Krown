-- +goose Up
CREATE TABLE scheduled_tournaments (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    entry_fee BIGINT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    recurrence_pattern TEXT NOT NULL,
    recurrence_start_timestamp TIMESTAMP NOT NULL,
    recurrence_end_timestamp TIMESTAMP NOT NULL,
    must_renew BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE scheduled_tournaments;