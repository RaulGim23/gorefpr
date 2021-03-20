package service

import (
	"context"

	"files/model"
)

// User godoc
type File interface {
	Find(ctx context.Context, id int64) (*model.File, error)
	FindAll(ctx context.Context) ([]model.File, error)
	Store(ctx context.Context, u *model.File) error
	Update(ctx context.Context, u *model.File) error
	Delete(ctx context.Context, id int64) error
}