-- name: CheckUserByEmail :one
SELECT COUNT(*) FROM users WHERE email = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (username, email, password) VALUES ($1, $2, $3);