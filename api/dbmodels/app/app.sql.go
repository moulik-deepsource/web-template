// Code generated by sqlc. DO NOT EDIT.
// source: app.sql

package app

import (
	"context"
)

const getByID = `-- name: GetByID :one
SELECT id, name, os_id, created_at, modified_at FROM app WHERE id = $1 LIMIT 1
`

func (q *Queries) GetByID(ctx context.Context, id int64) (App, error) {
	row := q.db.QueryRowContext(ctx, getByID, id)
	var i App
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OsID,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
