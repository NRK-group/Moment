package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestProfile(t *testing.T) {
	t.Run("Request with valid enpoint", func(t *testing.T) {
		// Create the database that will be used for testing

		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		w := httptest.NewRecorder()

		Env.Profile(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Request with invalid enpoint", func(t *testing.T) {
		// Create the database that will be used for testing
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
		w := httptest.NewRecorder()

		Env.Profile(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("No Cookie Present", func(t *testing.T) {
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		w := httptest.NewRecorder()

		Env.Profile(w, req)
		want := "Unauthorised"
		got := w.Body.String()
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Register invalid Cookie", func(t *testing.T) {
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		w := httptest.NewRecorder()
		req.AddCookie(&http.Cookie{Name: "session_token", Value: "INVALID"})

		Env.Profile(w, req)
		want := "Unauthorised"
		got := w.Body.String()
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Check correct user is returned", func(t *testing.T) {
		Env := handler.Env{Env: database}
		profileEmail := strings.ToLower("profile@" + uuid.NewV4().String() + ".com")
		// Create the struct that will be inserted
		sampleUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: profileEmail, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.Capitalise(&sampleUser)

		// Marhsal the struct to a slice of bytes
		sampleUserBytes, err := json.Marshal(sampleUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)

		req := httptest.NewRequest(http.MethodPost, "/registration", testReq)
		w := httptest.NewRecorder()
		Env.Registration(w, req)
		loginUser := structs.User{Email: profileEmail, Password: "TestPass"}

		// Marhsal the struct to a slice of bytes
		logUser, err := json.Marshal(loginUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		logReq := bytes.NewReader(logUser)
		reqLogin := httptest.NewRequest(http.MethodPost, "/login", logReq)
		wr := httptest.NewRecorder()
		Env.Login(wr, reqLogin)
		var temp structs.User
		getErr := auth.GetUser("email", profileEmail, &temp, *Env.Env)
		if getErr != nil {
			t.Errorf("Error getting the user")
			return
		}

		reqProf := httptest.NewRequest(http.MethodGet, "/profile", nil)
		reqProf.Header = http.Header{"Cookie": wr.Header()["Set-Cookie"]}

		values := reqProf.URL.Query()
		values.Add("userID", temp.UserId)
		reqProf.URL.RawQuery = values.Encode()

		profWr := httptest.NewRecorder()
		Env.Profile(profWr, reqProf)
		var result structs.User
		gotErr := json.Unmarshal(profWr.Body.Bytes(), &result)
		if gotErr != nil {
			t.Errorf("Error unmarshalling result: %v", gotErr)
			return
		}
		got := result
		want := sampleUser
		got.UserId = "-"
		got.SessionId = "-"
		got.CreatedAt = "-"
		want.Password = ""

		if got != want {
			t.Errorf("got %v \n want %v", got, want)
		}
	})
}
