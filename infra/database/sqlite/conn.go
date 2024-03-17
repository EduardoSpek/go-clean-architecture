package sqlite

import (
	"database/sql"
	"fmt"
	"os"
)

type SQLite struct {}

var conn SQLite

func (repo *SQLite) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("PATCH_DB_SQLITE"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}