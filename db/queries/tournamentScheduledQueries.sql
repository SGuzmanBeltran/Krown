-- name: GetScheduledTournaments :many
SELECT * FROM scheduled_tournaments;

-- name: GetScheduledTournament :one
SELECT * FROM scheduled_tournaments WHERE id = $1;

-- name: CreateScheduledTournament :one
INSERT INTO scheduled_tournaments (
    name,
    entry_fee,
    start_time,
    recurrence_pattern,
    recurrence_start_timestamp,
    recurrence_end_timestamp,
    must_renew
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetScheduledTournamentsByStartTime :many
SELECT * FROM scheduled_tournaments where start_time = $1;