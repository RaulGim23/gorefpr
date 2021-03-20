package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Database(mysql string) (*sql.DB, error) {
	db, err := sql.Open("mysql", mysql)
	if err != nil {
		return nil, err
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS "files" (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		file_name	VARCHAR(255),
		date	    VARCHAR(255)
	)`)

	return db, err
}