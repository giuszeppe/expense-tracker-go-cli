package main

import (
	"database/sql"
	"log"

	"github.com/giuszeppe/expense-tracker-go-cli/cmd"
	"github.com/giuszeppe/expense-tracker-go-cli/migrations"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	cmd.Execute()
}
