package file

import (
	"context"
	"errors"

	"files/model"
	"files/repository"
	"files/service"
)

type serviceImpl struct {
	File       repository.File
}

func New(u repository.File) service.File {
	return &serviceImpl{File: u}
}

func (s *serviceImpl) Find(ctx context.Context, id int64) (*model.File, error) {
	return s.File.Find(ctx, id)
}

// FindAll return a list of users
func (s *serviceImpl) FindAll(ctx context.Context) ([]model.File, error) {
	return s.File.FindAll(ctx)
}

// Store create a user
func (s *serviceImpl) Store(ctx context.Context, user *model.File) error {
	return s.File.Store(ctx, user)
}

// Update a user by ID
func (s *serviceImpl) Update(ctx context.Context, update *model.File) error {
	file, err := s.File.Find(ctx, update.ID)
	if err != nil {
		return errors.New("user not found ")
	}

	file.FileName = update.FileName
	file.Date = update.FileName
	return s.File.Update(ctx, file)
}

func (s *serviceImpl) Delete(ctx context.Context, id int64) error {
	return s.File.Delete(ctx, id)
}
