package main

import (
	"log"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/db/sqlite"
	"backend/pkg/follow"
	"backend/pkg/structs"
)

func main() {
	// this open or create the database
	networkDb := sqlite.CreateDatabase("./social_network.db")

	// migrate the database
	sqlite.MigrateDatabase("file://pkg/db/migrations/sqlite", "sqlite3://./social_network.db")

	// initialize the database struct
	data := &structs.DB{DB: networkDb}

	// close the database
	defer networkDb.Close()
	user1 := &structs.User{
		FirstName: "Ricky", LastName: "Founder", NickName: "adriell", Email: "adriell@gmail.com", Password: "Hello123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "images/profile/desktop.jpg", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err := auth.InsertUser(*user1, *data)
	if err != nil {
		log.Println(err)
	}
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *data)
	user2 := &structs.User{
		FirstName: "Nate", LastName: "Founder", NickName: "Nate", Email: "nate@gmail.com", Password: "Hello123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "images/profile/default-user.svg", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err = auth.InsertUser(*user2, *data)
	if err != nil {
		log.Println(err)
	}
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *data)
	follow.FollowUser(result1.UserId, result2.UserId, data)
	follow.FollowUser(result2.UserId, result1.UserId, data)
	chat.InsertNewChat(result1.UserId, result2.UserId, data)

	user3 := &structs.User{
		FirstName: "Kev", LastName: "Founder", NickName: "Kev", Email: "kev@gmail.com", Password: "Hello123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "images/profile/default-user.svg", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err = auth.InsertUser(*user3, *data)
	if err != nil {
		log.Println(user3.Email, err)
	}
	var result3 structs.User
	err = auth.GetUser("email", user3.Email, &result3, *data)
	if err != nil {
		log.Println(err)
	}
	follow.FollowUser(result1.UserId, result3.UserId, data)
	follow.FollowUser(result2.UserId, result3.UserId, data)
	follow.FollowUser(result3.UserId, result1.UserId, data)
	follow.FollowUser(result3.UserId, result2.UserId, data)
	chat.InsertNewChat(result1.UserId, result3.UserId, data)
	chat.InsertNewChat(result2.UserId, result3.UserId, data)
}
