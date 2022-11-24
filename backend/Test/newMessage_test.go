package Test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
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

func TestNewMessageHandler(t *testing.T) {
	Env := &handler.Env{Env: database}
	delete, _ := database.DB.Prepare(`DELETE FROM User`)
	delete.Exec()
	email1 := "test" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Ricky", LastName: "Founder", NickName: "adriell", Email: email1, Password: "Hello123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "images/profile/desktop.jpg", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err := auth.InsertUser(*user1, *Env.Env)
	if err != nil {
		log.Println(err)
	}
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *Env.Env)
	email2 := "test" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "Founder", NickName: "Nate", Email: email2, Password: "Hello123",
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
	sampleUser := &structs.User{
		Email: email1, Password: "Hello123",
	}
	sampleUserBytes, err := json.Marshal(sampleUser)
	t.Run("Marshal user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
	})
	testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
	req := httptest.NewRequest(http.MethodPost, "/login", testReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	t.Run("Login user", func(t *testing.T) {
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	req2 := httptest.NewRequest(http.MethodGet, "/message/new", nil)
	req2.Header.Set("Cookie", w.Header().Get("Set-Cookie"))
	w2 := httptest.NewRecorder()
	Env.NewMessage(w2, req2)
	t.Run("Get new messages", func(t *testing.T) {
		if w2.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w2.Code)
		}
		var result []structs.Info
		err := json.NewDecoder(w2.Body).Decode(&result)
		if err != nil {
			t.Errorf("Error decoding json: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Expected 1, got %v", len(result))
		}
		if result[0].Id != result2.UserId {
			t.Errorf("Expected %v, got %v", result2.UserId, result[0].Id)
		}
	})
}
