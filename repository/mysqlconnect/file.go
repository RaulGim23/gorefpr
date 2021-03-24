package mysqlconnect

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"

	"files/model"
	"files/repository"
)

const (
	sqlFileTable = "files"
)

type fileRepository struct {
	db repository.Database
}

// File the file repository.
func File(database repository.Database) repository.File {
	return &fileRepository{db: database}
}

// Find returns an file by ID.
func (r *fileRepository) Find(ctx context.Context, id int64) (*model.File, error) {
	return r.Scan(squirrel.Select("*").From(sqlFileTable).Where("id=?", id).
		RunWith(r.db).QueryRowContext(ctx))
}

// FindAll return a list of files.
func (r *fileRepository) FindAll(ctx context.Context) ([]model.File, error) {
	rows, err := squirrel.Select("*").From(sqlFileTable).RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []model.File

	for rows.Next() {
		file, err := r.Scan(rows)
		if err != nil {
			return nil, err
		}

		files = append(files, *file)
	}

	return files, nil
}

// Store create a file record.
func (r *fileRepository) Store(ctx context.Context, u *model.File) error {
	result, err := squirrel.Insert(sqlFileTable).Columns("file_name", "date").
		Values(u.FileName, u.Date).RunWith(r.db).ExecContext(ctx)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err
}

// Update a file record.
func (r *fileRepository) Update(ctx context.Context, f *model.File) error {
	sql := squirrel.Update(sqlFileTable)

	if strings.TrimSpace(f.FileName) != "" {
		sql = sql.Set("file_name", f.FileName)
	}

	if strings.TrimSpace(f.Date) != "" {
		sql = sql.Set("date", f.Date)
	}

	_, err := sql.Where("id=?", f.ID).RunWith(r.db).ExecContext(ctx)

	return err
}

// Delete a file by ID.
func (r *fileRepository) Delete(ctx context.Context, id int64) error {
	_, err := squirrel.Delete(sqlFileTable).Where("id=?", id).RunWith(r.db).ExecContext(ctx)

	return err
}

// Scan models.File row scanner.
func (r fileRepository) Scan(row repository.Scanner) (*model.File, error) {
	file := &model.File{}
	return file, row.Scan(
		&file.ID,
		&file.FileName,
		&file.Date,
	)
}
