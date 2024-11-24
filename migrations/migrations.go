package migrations

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE if not exists expenses(
		id integer PRIMARY KEY,
		amount float,
		description text,
		collected_date TEXT DEFAULT CURRENT_DATE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
