// Code generated by sqlc. DO NOT EDIT.
// source: os.sql

package os

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const countTotal = `-- name: CountTotal :one
SELECT COUNT(id) FROM os
`

func (q *Queries) CountTotal(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTotal)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOne = `-- name: CreateOne :one
INSERT INTO os(name, created_at, modified_at) VALUES($1, $2, $3) RETURNING id, name, created_at, modified_at
`

type CreateOneParams struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (q *Queries) CreateOne(ctx context.Context, arg CreateOneParams) (OS, error) {
	row := q.db.QueryRowContext(ctx, createOne, arg.Name, arg.CreatedAt, arg.ModifiedAt)
	var i OS
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getByID = `-- name: GetByID :one
SELECT id, name, created_at, modified_at FROM os WHERE id = $1 LIMIT 1
`

func (q *Queries) GetByID(ctx context.Context, id int64) (OS, error) {
	row := q.db.QueryRowContext(ctx, getByID, id)
	var i OS
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getManyByIDs = `-- name: GetManyByIDs :many
SELECT id, name, created_at, modified_at FROM os WHERE id = ANY($1::bigint[])
`

func (q *Queries) GetManyByIDs(ctx context.Context, dollar_1 []int64) ([]OS, error) {
	rows, err := q.db.QueryContext(ctx, getManyByIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OS
	for rows.Next() {
		var i OS
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
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

const listOffset = `-- name: ListOffset :many
SELECT id, name, created_at, modified_at FROM os LIMIT $1 OFFSET $2
`

type ListOffsetParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListOffset(ctx context.Context, arg ListOffsetParams) ([]OS, error) {
	rows, err := q.db.QueryContext(ctx, listOffset, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OS
	for rows.Next() {
		var i OS
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
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
