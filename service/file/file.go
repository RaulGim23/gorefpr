package file

import (
	"context"
	"errors"

	"files/model"
	"files/repository"
	"files/service"
)

type serviceImpl struct {
	File repository.File
}

// New returns the service implementation.
func New(u repository.File) service.File {
	return &serviceImpl{File: u}
}

// Find return a file by ID.
func (s *serviceImpl) Find(ctx context.Context, id int64) (*model.File, error) {
	return s.File.Find(ctx, id)
}

// FindAll return a list of files.
func (s *serviceImpl) FindAll(ctx context.Context, orderBys[]string, page, limit uint64) ([]model.File, uint64, error) {
	return s.File.FindAll(ctx, orderBys, page, limit)
}

// Store create a file.
func (s *serviceImpl) Store(ctx context.Context, user *model.File) error {
	return s.File.Store(ctx, user)
}

// Update a user by ID.
func (s *serviceImpl) Update(ctx context.Context, update *model.File) error {
	file, err := s.File.Find(ctx, update.ID)
	if err != nil {
		return errors.New("user not found ")
	}
	file.FileName = update.FileName
	file.Date = update.Date
	return s.File.Update(ctx, file)
}

// Delete deletes a file by ID.
func (s *serviceImpl) Delete(ctx context.Context, id int64) error {
	return s.File.Delete(ctx, id)
}
