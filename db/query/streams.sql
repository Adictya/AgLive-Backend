-- name: CreateStream :one
INSERT INTO streams (
	channel,
	thumbnail,
	streamer
) VALUES ($1,$2,$3)
RETURNING *;

-- name: ListStreams :many
SELECT id,channel,streamer FROM streams
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetThumbnail :one
SELECT thumbnail FROM streams WHERE id = $1;

-- name: DeleteStream :exec
DELETE FROM streams WHERE channel = $1 AND streamer = $2;
