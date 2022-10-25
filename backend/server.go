package main

import (
	"log"
	"net/http"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/structs"
)

func main() {
	// this open or create the database
	networkDb := sqlite.CreateDatabase("./social_network.db")

	// migrate the database
	sqlite.MigrateDatabase("file://pkg/db/migrations/sqlite", "sqlite3://./social_network.db")

	// initialize the database struct
	data := &structs.DB{DB: networkDb}
	database := &handler.Env{Env: data}

	// close the database
	defer networkDb.Close()

	// initialize the routes
	http.HandleFunc("/", database.Home) // set the handler for the home route

	// handler for the websocket
	hub := handler.NewHub()
	go hub.LogConns()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(hub, w, r)
	})

	// start the server
	log.Println("Server is running on port 5070")
	http.ListenAndServe(":5070", nil)
}
