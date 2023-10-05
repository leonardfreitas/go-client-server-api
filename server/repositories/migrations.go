package repositories

import (
	"database/sql"
)

func RunMigrations(db *sql.DB) error {
	stmt := `
		CREATE TABLE IF NOT EXISTS dollars(
			id INTEGER PRIMARY KEY, 
			code TEXT,
			codein TEXT,
			name TEXT,
			high TEXT,
			low TEXT,
			varbid TEXT,
			pct_change TEXT,
			bid TEXT,
			ask TEXT,
			timestamp TEXT,
			create_date TEXT
		);
	`
	_, err := db.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}
