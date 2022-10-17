package main

import (
	"fmt"
	"net/http"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
)

func main() {
	db := sqlite.CreateDatabase("./social_network.db")                                         // this open or create the database
	sqlite.MigrateDatabase("file://pkg/db/migrations/sqlite", "sqlite3://./social_network.db") // migrate the database
	database := &handler.DB{DB: db}                                                            // create a database structs
	http.HandleFunc("/", database.Home)
	defer db.Close() // close the database
	fmt.Println("Server is running on port 5070")
	http.ListenAndServe(":5070", nil)
	fmt.Println("setup")
}
