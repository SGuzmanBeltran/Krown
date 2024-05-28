-- +goose Up
CREATE TABLE users_teams (
    id BIGSERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    team_id INT NOT NULL,
    tournaments_played INT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (team_id) REFERENCES teams(id)
);

-- +goose Down
DROP TABLE users_teams;