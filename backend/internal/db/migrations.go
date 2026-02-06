package db

import (
	"database/sql"
	"slices"
)

type Migration struct {
	version int
	up      func(*sql.Tx) error
	down    func(*sql.Tx) error
}

var allMigrations = []Migration{}

const createTableMigrations = `
CREATE TABLE IF NOT EXISTS migrations (
	version INTEGER PRIMARY KEY,
	applied_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%ML%fZ', 'now')
);
`

func ensureMigrationsTable(db *sql.DB) error {
	_, err := db.Exec(createTableMigrations)
	return err
}

func getAppliedVersions(db *sql.DB) ([]int, error) {
	rawRows, err := db.Query("SELECT version FROM migrations ORDER BY ASC")
	if err != nil {
		return nil, err
	}
	defer rawRows.Close()

	var rows []int
	for rawRows.Next() {
		var version int
		err := rawRows.Scan(&version)
		if err != nil {
			return nil, err
		}
		rows = append(rows, version)
	}

	return rows, nil
}

func applyMigration(db *sql.DB, migration Migration) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = migration.up(tx)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO migrations (version) VALUES (?)", migration.version)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func RunMigrations(db *sql.DB) error {
	err := ensureMigrationsTable(db)
	if err != nil {
		return err
	}

	appliedVersions, err := getAppliedVersions(db)
	if err != nil {
		return err
	}

	for _, migration := range allMigrations {
		if !slices.Contains(appliedVersions, migration.version) {
			err = applyMigration(db, migration)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
