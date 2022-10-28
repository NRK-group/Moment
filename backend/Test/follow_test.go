package Test

import (
	"errors"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/follow"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// TestFollow will include all the tests for the follow feature
func TestFollow(t *testing.T) {
	// Create the database that will be used for testing
	database := &structs.DB{DB: sqlite.CreateDatabase("./social_network_test.db")}
	// migrate the database
	sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")
	followerId := uuid.NewV4().String()
	followingId := uuid.NewV4().String()
	// Create the users that will be used for testing
	email := logTestEmail
	inputUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	t.Run("Insert follow", func(t *testing.T) {
		got := follow.InsertFollow(followerId, followingId, database).Error()
		want := errors.New("no error").Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	followerId = "8a5e7163-7cb7-440a-9537-42a9c4371752"
	followingId = "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2"
	t.Run("Delete follow", func(t *testing.T) {
		got := follow.DeleteFollow(followerId, followingId, database).Error()
		want := errors.New("no error").Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Check follow false", func(t *testing.T) {
		got := follow.CheckIfFollow(followerId, followingId, database)
		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
	t.Run("Check follow true", func(t *testing.T) {
		follow.InsertFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		got := follow.CheckIfFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		want := true
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
	t.Run("Follow user returns follow", func(t *testing.T) {
		follow.DeleteFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		status, err := follow.FollowUser("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		got := err.Error()
		want := errors.New("no error").Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		got = status
		want = "follow"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Follow user returns pending", func(t *testing.T) {
		if !follow.CheckIfFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database) {
			follow.InsertFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		}
		status, err := follow.FollowUser("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		got := err.Error()
		want := errors.New("no error").Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		got = status
		want = "pending"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Follow user returns follow", func(t *testing.T) {
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		got := ""
		want := "follow"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
