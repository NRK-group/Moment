package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/follow"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestChat(t *testing.T) {
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}

	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user2, *database)
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *database)
	follow.FollowUser(result1.UserId, result2.UserId, database)
	follow.FollowUser(result2.UserId, result1.UserId, database)
	t.Run("Insert new chat", func(t *testing.T) {
		_, err := chat.InsertNewChat(result1.UserId, result2.UserId, database)
		if err != nil {
			t.Errorf("InsertNewChat() error = %v", err)
		}
	})
	t.Run("Get Previous Private chat", func(t *testing.T) {
		chats, err := chat.GetPreviousPrivateChat(result1.UserId, database)
		if err != nil {
			t.Errorf("GetPreviousPrivateChat() error = %v", err)
		}
		if len(chats) != 1 {
			t.Errorf("Error expected 1 chat, got %v", len(chats))
		}
	})
	t.Run("Insert New Group Chat", func(t *testing.T) {
		err := chat.InsertNewGroupChat("hello", database)
		if err != nil {
			t.Errorf("InsertNewGroupChat() error = %v", err)
		}
	})
}
