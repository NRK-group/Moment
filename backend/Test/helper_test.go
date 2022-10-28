package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/helper"
	"backend/pkg/structs"
)

func TestHelper(t *testing.T) {
	database := &structs.DB{DB: sqlite.CreateDatabase("./social_network_test.db")}
	sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")
	email := logTestEmail
	inputUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*inputUser, *database)
	var result structs.User
	auth.GetUser("email", email, &result, *database)
	t.Run("Check user privacy", func(t *testing.T) {
		got := helper.CheckUserPrivacy(result.UserId, database)
		want := "public"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
