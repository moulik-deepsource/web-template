// Code generated by sqlc. DO NOT EDIT.

package os

import (
	"context"
)

type Querier interface {
	CountTotal(ctx context.Context) (int64, error)
	CreateOne(ctx context.Context, arg CreateOneParams) (OS, error)
	GetByID(ctx context.Context, id int64) (OS, error)
	GetManyByIDs(ctx context.Context, dollar_1 []int64) ([]OS, error)
	ListOffset(ctx context.Context, arg ListOffsetParams) ([]OS, error)
}

var _ Querier = (*Queries)(nil)
