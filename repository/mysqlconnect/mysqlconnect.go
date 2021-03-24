package mysqlconnect

import (
	"database/sql"
)

// Database godoc.
func Database(mysql string) (*sql.DB, error) {
	db, err := sql.Open("mysql", mysql)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "files" (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		file_name	VARCHAR(255),
		date	    VARCHAR(255)
	)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
