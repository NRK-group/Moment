package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestNotification(t *testing.T) {
	Env := handler.Env{Env: database}
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	sampleUser := &structs.User{
		Email: email, Password: "Password123",
	}
	sampleUserBytes, err := json.Marshal(sampleUser)
	t.Run("Test Marshal data", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
	})
	testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
	req := httptest.NewRequest(http.MethodPost, "/login", testReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	t.Run("Test login", func(t *testing.T) {
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	req2 := httptest.NewRequest(http.MethodGet, "/notification", nil)
	req2.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values := req2.URL.Query()
	values.Add("notifType", "follow")
	req2.URL.RawQuery = values.Encode()
	w2 := httptest.NewRecorder()
	Env.Notification(w2, req2)
	t.Run("Test notification", func(t *testing.T) {
		if w2.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w2.Code)
		}
		if w2.Body.String() != "follow" {
			t.Errorf("Expected body %s, got %s", "follow", w2.Body.String())
		}
	})
	req3 := httptest.NewRequest(http.MethodGet, "/notification", nil)
	req3.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values = req3.URL.Query()
	values.Add("notifType", "general")
	req3.URL.RawQuery = values.Encode()
	w3 := httptest.NewRecorder()
	Env.Notification(w3, req3)
	t.Run("Test notification", func(t *testing.T) {
		if w3.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w3.Code)
		}
		if w3.Body.String() != "general" {
			t.Errorf("Expected body %s, got %s", "general", w3.Body.String())
		}
	})
	req4 := httptest.NewRequest(http.MethodGet, "/notification", nil)
	req4.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values = req4.URL.Query()
	values.Add("notifType", "group")
	req4.URL.RawQuery = values.Encode()
	w4 := httptest.NewRecorder()
	Env.Notification(w4, req4)
	t.Run("Test notification", func(t *testing.T) {
		if w4.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w4.Code)
		}
		if w4.Body.String() != "group" {
			t.Errorf("Expected body %s, got %s", "group", w4.Body.String())
		}
	})
}
