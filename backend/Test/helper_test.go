package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/helper"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestHelper(t *testing.T) {
	// database := DatabaseSetup()
	email := "Helper" + uuid.NewV4().String() + "@test.com"
	inputUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*inputUser, *database)
	var result structs.User
	auth.GetUser("email", email, &result, *database)
	t.Run("Check user privacy", func(t *testing.T) {
		got := helper.CheckUserIfPublic(result.UserId, database)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
