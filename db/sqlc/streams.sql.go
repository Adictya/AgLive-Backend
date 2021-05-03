// Code generated by sqlc. DO NOT EDIT.
// source: streams.sql

package db

import (
	"context"
)

const createStream = `-- name: CreateStream :one
INSERT INTO streams (
	channel,
	thumbnail
) VALUES ($1,$2)
RETURNING id, channel, thumbnail, streamer
`

type CreateStreamParams struct {
	Channel   string `json:"channel"`
	Thumbnail string `json:"thumbnail"`
}

func (q *Queries) CreateStream(ctx context.Context, arg CreateStreamParams) (Stream, error) {
	row := q.db.QueryRowContext(ctx, createStream, arg.Channel, arg.Thumbnail)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.Channel,
		&i.Thumbnail,
		&i.Streamer,
	)
	return i, err
}

const deleteStream = `-- name: DeleteStream :exec
DELETE FROM streams WHERE channel = $1
`

func (q *Queries) DeleteStream(ctx context.Context, channel string) error {
	_, err := q.db.ExecContext(ctx, deleteStream, channel)
	return err
}

const getThumbnail = `-- name: GetThumbnail :one
SELECT thumbnail FROM streams WHERE id = $1
`

func (q *Queries) GetThumbnail(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, getThumbnail, id)
	var thumbnail string
	err := row.Scan(&thumbnail)
	return thumbnail, err
}

const listStreams = `-- name: ListStreams :many
SELECT id,channel FROM streams
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStreamsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListStreamsRow struct {
	ID      int32  `json:"id"`
	Channel string `json:"channel"`
}

func (q *Queries) ListStreams(ctx context.Context, arg ListStreamsParams) ([]ListStreamsRow, error) {
	rows, err := q.db.QueryContext(ctx, listStreams, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListStreamsRow
	for rows.Next() {
		var i ListStreamsRow
		if err := rows.Scan(&i.ID, &i.Channel); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
