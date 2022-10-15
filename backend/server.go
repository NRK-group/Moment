package main

import (
	"fmt"

	"backend/pkg/db/sqlite"
)

func main() {
	db := sqlite.CreateDatabase() // this open or create the database
	sqlite.MigrateDatabase()      // migrate the database
	db.Close()                    // close the database
	fmt.Println("setup")
}
