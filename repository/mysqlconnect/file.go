package mysqlconnect

import (
	"context"
	"fmt"
	"strings"

	"files/model"
	"files/repository"
)

const (
	sqlGetUser		= "SELECT * FROM files WHERE id=?"
	sqlSelectFiles 	= "SELECT * FROM files"
	sqlInsertFile 	= "INSERT INTO files (file_name, date) VALUES(?, ?)"
	sqlUpdateFile 	= "UPDATE files SET (%s) WHERE id=?"
	sqlDeleteFile 	= "DELETE FROM files WHERE id=?"
)

type fileRepository struct {
	db repository.Database
}

func File(database repository.Database) repository.File {
	return &fileRepository{db: database}
}

// Find returns an user
func (r *fileRepository) Find(ctx context.Context, id int64) (*model.File, error) {
	return r.Scan(r.db.QueryRowContext(ctx, sqlGetUser, id))
}

// FindAll
func (r *fileRepository) FindAll(ctx context.Context) ([]model.File, error) {
	rows, err := r.db.QueryContext(ctx, sqlSelectFiles)
	if err != nil {
		return nil, err;
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

// Store
func (r *fileRepository) Store(ctx context.Context, u *model.File) error {
	result, err := r.db.ExecContext(ctx, sqlInsertFile, u.FileName, u.Date)
	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	return err
}

// Update a file
func (r *fileRepository) Update(ctx context.Context, u *model.File) error {
	var fields []string
	var params []interface{}
	if strings.TrimSpace(u.FileName) != "" {
		fields = append(fields, "`file_name`=?")
		params = append(params, u.FileName)
	}
	if strings.TrimSpace(u.Date) != "" {
		fields = append(fields, "`date`=?")
		params = append(params, u.Date)
	}

	query := fmt.Sprintf(sqlUpdateFile, strings.Join(fields, ","))
	_, err := r.db.ExecContext(ctx, query, append(params, u.ID)...)
	return err
}

// Delete a user by ID
func (r *fileRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, sqlDeleteFile, id)
	return err
}

// Scan models.User row scanner
func (r *fileRepository) Scan(row repository.Scanner) (*model.File, error) {
	file := &model.File{}
	return file, row.Scan(
		&file.ID,
		&file.FileName,
		&file.Date,
	)
}
