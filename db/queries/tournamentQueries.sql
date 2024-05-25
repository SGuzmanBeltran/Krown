-- name: GetTournaments :many
SELECT * FROM tournaments WHERE start_time >= $1 and start_time <= $2;

-- name: GetTournament :one
SELECT * FROM tournaments WHERE id = $1;

-- name: BatchCreate :batchone
INSERT INTO tournaments (
    name,
    entry_fee,
    start_time
) VALUES (
    $1, $2, $3
) RETURNING *;