package main

import (
	"log"
	"net/http"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
)

func main() {
	// this open or create the database
	db := sqlite.CreateDatabase("./social_network.db")

	// migrate the database
	sqlite.MigrateDatabase("file://pkg/db/migrations/sqlite", "sqlite3://./social_network.db")

	// initialize the database struct
	database := &handler.DB{DB: db}

	// close the database
	defer db.Close()

	// initialize the routes
	http.HandleFunc("/", database.Home) // set the handler for the home route
	http.HandleFunc("/ws", handler.WsEndpoint)

	// start the server
	log.Println("Server is running on port 5070")
	http.ListenAndServe(":5070", nil)
}
