package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func CreateDatabase() {
	db, err := sql.Open("sqlite3", "./social_network.db")
	if err != nil {
		log.Fatal("Database open error ", err)
	}
	fmt.Println("Database created", db.Stats())
	// stmt, err := db.Prepare(`CREATE DATABASE social_network`)
	// if err != nil {
	// 	log.Fatal("Database prepare error ", err)
	// }
	// _, err = stmt.Exec()
	// if err != nil {
	// 	log.Fatal("Database exec error: ", err)
	// }
	defer db.Close()
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./pkg/db/migrations/sqlite")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}
	// m, err := migrate.New(
	// 	"file://backend/pkg/db/sqlite/migrations",
	// 	"sqlite3://./social_network.db")
	if err != nil {
		log.Fatal(err)
	}

	// modify for Down
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

// func MigrateDatabase() {
// 	m, err := migrate.New(
// 		"file://backend/pkg/db/sqlite/migrations",
// 		"sqlite3://./social_network.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = m.Steps(1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
