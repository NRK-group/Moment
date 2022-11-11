package main

import (
	"log"
	"net/http"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/structs"
	wSocket "backend/pkg/websocket"
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
	SetUpRoutes(database)

	gallery := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", gallery)) // handling the CSS

	// handler for the websocket
	hub := wSocket.NewHub()
	go hub.LogConns()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		database.ServeWs(hub, w, r)
	})

	// start the server
	log.Println("Server is running on port 5070")
	http.ListenAndServe(":5070", nil)
}

// SetUpRoutes initialises the handlers
func SetUpRoutes(database *handler.Env) {
	http.HandleFunc("/", database.Home)
	http.HandleFunc("/post", database.Post)
	http.HandleFunc("/group", database.Group)
	http.HandleFunc("/event", database.Event)
	http.HandleFunc("/login", database.Login)
	http.HandleFunc("/logout", database.Logout)
	http.HandleFunc("/registration", database.Registration)
	http.HandleFunc("/validate", database.Validate)
	http.HandleFunc("/updateprofileimg", database.UpdateImage)
	http.HandleFunc("/chat", database.Chat)
	http.HandleFunc("/profile", database.Profile)
	http.HandleFunc("/message", database.Message)
}
