package Test

import (
	"log"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestNewMessage(t *testing.T) {
	Env := &handler.Env{Env: database}
	email := "test" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Ricky", LastName: "Founder", NickName: "adriell", Email: email, Password: "Hello123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "images/profile/desktop.jpg", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err := auth.InsertUser(*user1, *Env.Env)
	if err != nil {
		log.Println(err)
	}
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *Env.Env)
	email = "test" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "Founder", NickName: "Nate", Email: email, Password: "Hello123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "images/profile/default-user.svg", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err = auth.InsertUser(*user2, *Env.Env)
	if err != nil {
		log.Println(err)
	}
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *Env.Env)
	follow.FollowUser(result1.UserId, result2.UserId, Env.Env)
	follow.FollowUser(result2.UserId, result1.UserId, Env.Env)
	t.Run("Get user followers", func(t *testing.T) {
		follower, err := follow.GetFollowers(result1.UserId, Env.Env)
		if err != nil {
			t.Errorf("Error getting followers: %v", err)
		}
		if follower[0].FollowerId != result2.UserId {
			t.Errorf("Expected %v, got %v", result2.UserId, follower[0].FollowerId)
		}
		if len(follower) != 1 {
			t.Errorf("Error getting followers: %v", err)
		}
	})
	t.Run("Get user following", func(t *testing.T) {
		following, err := follow.GetFollowing(result1.UserId, Env.Env)
		if err != nil {
			t.Errorf("Error getting following: %v", err)
		}
		if following[0].FollowingId != result2.UserId {
			t.Errorf("Expected %v, got %v", result2.UserId, following[0].FollowingId)
		}
		if len(following) != 1 {
			t.Errorf("Error getting following: %v", err)
		}
	})
}
