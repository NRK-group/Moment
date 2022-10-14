package main

import (
	"fmt"

	"backend/pkg/db/sqlite"
	// _ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// _ "github.com/golang-migrate/migrate/v4/source/github"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlite.CreateDatabase()
	// sqlite.MigrateDatabase()
	fmt.Println("setup")
}
