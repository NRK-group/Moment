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
	wSocket "backend/pkg/websocket"

	uuid "github.com/satori/go.uuid"
)

var postIdarr []string

func twentyPost(database *structs.DB, user structs.User) {
	for i := 0; i < 10; i++ {
		post1 := structs.Post{UserID: user.UserId, Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry " + fmt.Sprint(i) + ".", GroupID: "3233234" + fmt.Sprint(i%2), Image: "https://picsum.photos/400/500"}
		str, _ := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, database)
		postIdarr = append([]string{str}, postIdarr...)
		fmt.Println(str)
	}
}

func randomComments(database *structs.DB, user structs.User) {
	// sampleUser := &structs.User{
	//  FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: "randEmail@gmail.com", Password: "InsertUser",
	//  DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
	//  IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	// }
	// auth.InsertUser(*sampleUser, *database)

	content := "test"

	for i := 0; i < 50; i++ {
		id, err := commets.CreateComment(user.UserId, postIdarr[rand.Intn(9)], content+fmt.Sprint(i), database)
		fmt.Println("err", err, "id", id)
	}
}

func CreateUser(database *structs.DB) structs.User {
	currentEmail := "hello" + uuid.NewV4().String() + "@test.com"
	currentUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: currentEmail, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*currentUser, *database)
	var currentResult structs.User
	auth.GetUser("email", currentEmail, &currentResult, *database)
	return currentResult
}

func main() {
	// this open or create the database
	networkDb := sqlite.CreateDatabase("./social_network.db")

	// migrate the database
	sqlite.MigrateDatabase("file://pkg/db/migrations/sqlite", "sqlite3://./social_network.db")

	// initialize the database struct
	data := &structs.DB{DB: networkDb}
	database := &handler.Env{Env: data}

	// newUser := CreateUser(data)
	// twentyPost(data, newUser)
	// randomComments(data, newUser)

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
	http.HandleFunc("/comment/", database.Comment)
	http.HandleFunc("/profile", database.Profile)
	http.HandleFunc("/message", database.Message)
	http.HandleFunc("/following", database.Following)
	http.HandleFunc("/followrequest", database.FollowReq)
	http.HandleFunc("/message/new", database.NewMessage)
	http.HandleFunc("/imageUpload", database.ImageUpload)
}
