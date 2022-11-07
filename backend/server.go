package main

import (

	"backend/pkg/commets"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/post"
	"backend/pkg/structs"
	wSocket "backend/pkg/websocket"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var postIdarr []string

func twentyPost(database *structs.DB) {
	for i := 0; i < 20; i++ {
		post1 := structs.Post{UserID: "3232131221" + fmt.Sprint(i), Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry " + fmt.Sprint(i) + ".", GroupID: "3233234" + fmt.Sprint(i%2), Image: "https://picsum.photos/400/500"}
		str, _ := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, database)
		postIdarr = append([]string{str}, postIdarr...)
		fmt.Println(str)
	}
}

func randomComments(database *structs.DB) {
	

	// sampleUser := &structs.User{
	// 	FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: "randEmail@gmail.com", Password: "InsertUser",
	// 	DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
	// 	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	// }
	// auth.InsertUser(*sampleUser, *database)

	content := "test"

	for i := 0; i < 50; i++ {
		commets.CreateComment("-", postIdarr[rand.Intn(19)], content+fmt.Sprint(i), database)
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

	//twentyPost(data)
	//randomComments(data)
	
	// initialize the routes
	http.HandleFunc("/", database.Home)
	http.HandleFunc("/post", database.Post)
	http.HandleFunc("/group", database.Group)
	http.HandleFunc("/event", database.Event)
	http.HandleFunc("/login", database.Login)
	http.HandleFunc("/comment", database.Comment)
	http.HandleFunc("/registration", database.Registration)
	http.HandleFunc("/validate", database.Validate)



	// handler for the websocket
	hub := wSocket.NewHub()
	go hub.LogConns()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(hub, w, r)
	})

	// start the server
	log.Println("Server is running on port 5070")
	http.ListenAndServe(":5070", nil)
}
