-- name: CreateStream :one
INSERT INTO streams (
	channel,
	thumbnail
) VALUES ($1,$2)
RETURNING *;

-- name: ListStreams :many
SELECT id,channel FROM streams
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetThumbnail :one
SELECT thumbnail FROM streams WHERE id = $1;


-- name: DeleteStream :exec
DELETE FROM streams WHERE channel = $1;
