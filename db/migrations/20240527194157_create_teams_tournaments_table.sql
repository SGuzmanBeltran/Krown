-- +goose Up
CREATE TABLE teams_tournaments (
    id BIGSERIAL PRIMARY KEY,
    team_id INT NOT NULL,
    tournament_id INT NOT NULL,
    enroll_payed BOOLEAN DEFAULT FALSE,
    team_position INT,
    tournaments_played INT DEFAULT 0,
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id)
);

-- +goose Down
DROP TABLE teams_tournaments;