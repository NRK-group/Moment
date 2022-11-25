package main

import (
	"log"
	"net/http"
	"os"

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
	// CreateingTengroups(data)
	// close the database
	defer networkDb.Close()

	// initialize the routes
	SetUpRoutes(database)

	gallery := http.FileServer(http.Dir("./images"))

	http.Handle("/images/", http.StripPrefix("/images/", gallery)) // handling the CSS

	// handler for the websocket
	hub := wSocket.NewHub(data)
	go hub.LogConns()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		database.ServeWs(hub, w, r)
	})

	// start the server
	port := ""
	port = os.Getenv("PORT")
	if port == "" {
		port = "8800"
	}
	log.Print("Listening on 0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
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
	http.HandleFunc("/comment/", database.Comment)
	http.HandleFunc("/profile", database.Profile)
	http.HandleFunc("/updateprofileinfo", database.ProfileChange)
	http.HandleFunc("/message", database.Message)
	http.HandleFunc("/following", database.Following)
	http.HandleFunc("/getfollowing", database.GetFollowing)
	http.HandleFunc("/followers", database.Followers)
	http.HandleFunc("/closefriend", database.CloseFriends)
	http.HandleFunc("/getclosefriend", database.CloseFriendList)
	http.HandleFunc("/followrequest", database.FollowReq)
	http.HandleFunc("/notification", database.Notification)
	http.HandleFunc("/message/new", database.NewMessage)
	http.HandleFunc("/imageUpload", database.ImageUpload)
	http.HandleFunc("/search", database.Search)
	http.HandleFunc("/getUserPosts", database.GetUserPosts)
	http.HandleFunc("/getGroupPost", database.GetGroupPost)
	http.HandleFunc("/getUserGroups", database.GetUserGroups)
	http.HandleFunc("/updateEventParticipant", database.UpdateEventParticipant)
	http.HandleFunc("/groupNonMembers", database.GroupNonMembers)
}
