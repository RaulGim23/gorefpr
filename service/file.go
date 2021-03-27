package service

import (
	"context"

	"files/model"
)

// File service definition.
type File interface {
	Find(ctx context.Context, id int64) (*model.File, error)
	FindAll(ctx context.Context, orderBys []string, page, limit uint64) ([]model.File, uint64, error)
	Store(ctx context.Context, u *model.File) error
	Update(ctx context.Context, u *model.File) error
	Delete(ctx context.Context, id int64) error
}
