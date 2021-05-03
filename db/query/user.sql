-- name: CreateUser :one
INSERT INTO users (
	username,
	hashed_password
) VALUES ($1,$2)
RETURNING *;

-- name: GerUser :one
SELECT * FROM users
WHERE username=$1 LIMIT 1;
