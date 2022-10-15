package sqlite

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func LogError(filename, funcName string, err error) {
	if err != nil {
		log.Printf("Error in %s/%s: %s\n\n", filename, funcName, err)
	}
}

func CreateDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./social_network.db") // this open or create the database
	if err != nil {
		LogError("sqlite.go", "CreateDatabase", err)
	}
	return db
}

func MigrateDatabase() {
	m, err := migrate.New(
		"file://pkg/db/migrations/sqlite/",
		"sqlite3://./social_network.db")
	if err != nil {
		LogError("sqlite.go", "MigrateDatabase", err)
	}
	err = m.Steps(1) // 1 is the step number to migrate up from the current version
	if err != nil {
		LogError("sqlite.go", "MigrateDatabase", err)
	}
}
