-- +goose Up
CREATE TABLE teams (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id INT NOT NULL,
    players_count INT DEFAULT 0,
    tournaments_played INT DEFAULT 0,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE teams;