package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/follow"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestPrivateChat(t *testing.T) {
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}

	auth.InsertUser(*user1, *database)
	var result1 structs.User
	err := auth.GetUser("email", user1.Email, &result1, *database)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	err = auth.InsertUser(*user2, *database)
	t.Run("Insert user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error inserting test struct %v", err)
		}
	})
	var result2 structs.User
	err = auth.GetUser("email", user2.Email, &result2, *database)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	_, err = follow.FollowUser(result1.UserId, result2.UserId, database)
	t.Run("Follow user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error following user %v", err)
		}
	})
	_, err = follow.FollowUser(result2.UserId, result1.UserId, database)
	t.Run("Follow user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error following user %v", err)
		}
	})
	_, err = chat.InsertNewChat(result1.UserId, result2.UserId, database)
	t.Run("Insert new chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("InsertNewChat() error = %v", err)
		}
	})
	chats, err := chat.GetPreviousPrivateChat(result1.UserId, database)
	t.Run("Get Previous Private chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("GetPreviousPrivateChat() error = %v", err)
		}
	})
	t.Run("Get Previous Private chat", func(t *testing.T) {
		if len(chats) != 1 {
			t.Errorf("Error expected 1 chat, got %v", len(chats))
		}
	})
	err = chat.InsertNewGroupChat("hello", database)
	t.Run("Insert New Group Chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("InsertNewGroupChat() error = %v", err)
		}
	})
	_, err = chat.GetUserInfo(result1.UserId, database)
	t.Run("Get userinfo for the chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("GetUserInfoForChat() error = %v", err)
		}
	})
}

func TestChatHandler(t *testing.T) {
	Env := handler.Env{Env: database}
	validationValidEmail := "add" + uuid.NewV4().String() + "@test.com"
	inputUser := &structs.User{
		FirstName: "Adriell", LastName: "Validation123", NickName: "", Email: validationValidEmail, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	t.Run("Chat handler", func(t *testing.T) {
		err := auth.InsertUser(*inputUser, *Env.Env)
		t.Run("Insert new user", func(t *testing.T) {
			if err != nil {
				t.Errorf("Error inserting test struct %v", err)
			}
		})
		var result structs.User
		err = auth.GetUser("email", validationValidEmail, &result, *Env.Env)
		t.Run("Get user", func(t *testing.T) {
			if err != nil {
				t.Errorf("Error getting test struct %v", err)
			}
		})
		sampleUser := &structs.User{
			Email: validationValidEmail, Password: "Password123",
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
		nr := httptest.NewRequest(http.MethodGet, "/chat", nil) // create a request
		nr.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
		nrr := httptest.NewRecorder() // create a response recorder
		Env.Chat(nrr, nr)             // call the handler
		// get cookies
		cookie, _ := nr.Cookie("session_token")
		got, _ := auth.SliceCookie(cookie.Value)
		want := result.UserId
		t.Run("Get cookie", func(t *testing.T) {
			if got[0] != want {
				t.Errorf("Expected %s, got %s", want, got[0])
			}
		})
		t.Run("Get chat", func(t *testing.T) {
			if nrr.Code != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, nrr.Code)
			}
		})
	})
}
