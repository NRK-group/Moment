package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/post"
	"backend/pkg/structs"
)


func twentyPost(database *structs.DB) {
	for i :=0 ; i < 20; i++ {
		post1 := structs.Post{UserID: "3232131221"+fmt.Sprint(i), Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry "+fmt.Sprint(i)+".", GroupID: "3233234"+fmt.Sprint(i % 2), Image: "https://picsum.photos/400/500"}
		str, _ := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, database)
		fmt.Println(str)
	}
}


func randomComments(database *structs.DB) {
	for i :=0 ; i < 20; i++ {

	}
}


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
	twentyPost(data)
	// initialize the routes
	http.HandleFunc("/", database.Home) 
	http.HandleFunc("/post", database.Post) 
	http.HandleFunc("/group", database.Group) 
	http.HandleFunc("/event", database.Event) 
	http.HandleFunc("/login", database.Login) 


	// start the server
	log.Println("Server is running on port 5070")
	http.ListenAndServe(":5070", nil)
}


