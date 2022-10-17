package main

import (
	"fmt"

	"backend/pkg/db/sqlite"
)

func main() {
	db := sqlite.CreateDatabase("./social_network.db")                                         // this open or create the database
	sqlite.MigrateDatabase("file://pkg/db/migrations/sqlite", "sqlite3://./social_network.db") // migrate the database
	defer db.Close()                                                                           // close the database
	fmt.Println("setup")
}
