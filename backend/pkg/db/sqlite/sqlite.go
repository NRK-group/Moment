package sqlite

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// LogError logs the error
//
//	filename: the file name where the error occurred
//	funcName: the function name where the error occurred
//	err: the error
func LogError(filename, funcName string, err error) {
	if err != nil {
		log.Printf("Error in %s/%s: %s\n\n", filename, funcName, err)
	}
}

// CreateDatabase opens or creates the database
//
// filename: the path and file name where the database is located
//
//	example: "./social_network.db"
func CreateDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path) // this open or create the database
	if err != nil {
		LogError("sqlite.go", "CreateDatabase", err)
	}
	return db
}

// MigrateDatabase migrates the database
//
// soureURL: the url where the migrations are located
//
//	ex: "file://./pkg/db/migrations/sqlite"
//
// databaseURL: the url where the database is located
//
//	ex: "sqlite3://./social_network.db"
//
// steps: the number of steps to migrate
//
//	ex: 1 (migrate 1 step from the current version) or -1 (migrate 1 step back from the current version)
func MigrateDatabase(soureURL, databaseURL string, steps int) {
	m, err := migrate.New(soureURL, databaseURL)
	if err != nil {
		LogError("sqlite.go", "MigrateDatabase", err)
	}
	// m.Up() // migrate up to the latest version
	// m.Down() // migrate down to the first version
	// m.Steps(-1) // migrate down by one version
	err = m.Steps(steps) // 1 is the step number to migrate up from the current version
	if err != nil {
		LogError("sqlite.go", "MigrateDatabase", err)
	}
}
