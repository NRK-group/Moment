package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/commets"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/post"
	"backend/pkg/structs"
)

func twentyPost(database *structs.DB) {
	for i := 0; i < 20; i++ {
		post1 := structs.Post{UserID: "3232131221" + fmt.Sprint(i), Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry " + fmt.Sprint(i) + ".", GroupID: "3233234" + fmt.Sprint(i%2), Image: "https://picsum.photos/400/500"}
		str, _ := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, database)
		fmt.Println(str)
	}
}

func randomComments(database *structs.DB) {
	postIdarr := [9]string{
		"aecb7625-8693-41d8-8f21-10a5bc391d04", "a9583a35-ed91-4750-b8cc-c1a870e0af49",
		"4f69494e-05b4-4608-a549-dca68c134363", "daf64d42-a466-433a-8645-4c2a5956273f",
		"c2dccbb2-38f0-47e5-aa57-27ffaac20ecf", "10874a69-0a80-405c-bbd1-05f157383ecf",
		"710d9ef6-9155-4ff8-9665-508f8a5f12ae", "78a11b20-2636-4a6f-8876-291acf10b449",
		"c604a90d-1ca9-48b1-8cee-3436a664e4cd",
	}

	sampleUser := &structs.User{
		FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: "randEmail@gmail.com", Password: "InsertUser",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
      auth.InsertUser(*sampleUser, *database)

	content := "test"

	for i := 0; i < 30; i++ {
		commets.CreateComment("-", postIdarr[rand.Intn(9)], content+fmt.Sprint(i), database)
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
	//randomComments(data)
	// twentyPost(data)
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
