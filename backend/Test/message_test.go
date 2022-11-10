package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/follow"
	l "backend/pkg/log"
	"backend/pkg/messages"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestMessage(t *testing.T) {
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}

	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user2, *database)
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *database)
	follow.FollowUser(result1.UserId, result2.UserId, database)
	follow.FollowUser(result2.UserId, result1.UserId, database)
	chatId, _ := chat.InsertNewChat(result1.UserId, result2.UserId, database)
	message := structs.Message{
		ChatId:     chatId,
		SenderId:   result1.UserId,
		ReceiverId: result2.UserId,
		Content:    "Hello World",
	}
	t.Run("Test Insert Message", func(t *testing.T) {
		msg, err := messages.InsertMessage(message, *database)
		l.LogMessage("message_test.go", "Test Insert Message", msg)
		if err != nil {
			t.Errorf("Error inserting message: %v", err)
		}
	})
}
