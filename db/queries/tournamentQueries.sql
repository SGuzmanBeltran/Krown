-- name: GetTournaments :many
SELECT * FROM tournaments WHERE start_time > $1 and start_time < $2;

-- name: GetTournament :one
SELECT * FROM tournaments WHERE id = $1;