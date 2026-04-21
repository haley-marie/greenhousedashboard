package db

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func OpenDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	row := db.QueryRow("PRAGMA foreign_keys;")
	var fk int
	row.Scan(&fk)

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	_, err = db.Exec("PRAGMA busy_timeout = 5000")
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
