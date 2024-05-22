-- name: GetScheduledTournaments :many
SELECT * FROM scheduled_tournaments;

-- name: GetScheduledTournament :one
SELECT * FROM scheduled_tournaments WHERE id = $1;