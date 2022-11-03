package sqlite

import (
	"database/sql"

	l "backend/pkg/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// CreateDatabase opens or creates the database
//
// filename: the path and file name where the database is located
//
//	example: "./social_network.db"
func CreateDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path) // this open or create the database
	if err != nil {
		l.LogMessage("sqlite.go", "CreateDatabase", err)
	}
	return db
}

// MigrateDatabase migrates the database to the latest version
//
// soureURL: the url where the migrations are located
//
//	ex: "file://./pkg/db/migrations/sqlite"
//
// databaseURL: the url where the database is located
//
//	ex: "sqlite3://./social_network.db"
func MigrateDatabase(soureURL, databaseURL string) {
	m, err := migrate.New(soureURL, databaseURL)
	if err != nil {
		l.LogMessage("sqlite.go", "MigrateDatabase", err)
	}
	// m.Up() // migrate up to the latest version
	// m.Down() // migrate down to the first version
	// m.Steps(-1) // migrate down by one version
	err = m.Up() // 1 is the step number to migrate up from the current version
	if err != nil {
		l.LogMessage("sqlite.go", "MigrateDatabase", err)
	}
}
