package repository

import (
	"context"
	"database/sql"
)

// Database represent sql database minimal interface
type Database interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// Scanner
type Scanner interface {
	Scan(dest ...interface{}) error
}