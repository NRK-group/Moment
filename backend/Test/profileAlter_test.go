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

func TestProfileUpdate(t *testing.T) {
	t.Run("update with valid user", func(t *testing.T) {
		// register a user
		Env := handler.Env{Env: database}
		handlerEmail := "handlertest@" + uuid.NewV4().String() + ".com"
		repEmail := "new@" + uuid.NewV4().String() + ".com"

		// Create the struct that will be inserted
		sampleUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: handlerEmail, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

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
		loginUser := structs.User{Email: sampleUser.Email, Password: sampleUser.Password}
		// Marhsal the struct to a slice of bytes
		loginBytes, err := json.Marshal(loginUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		loginReq := bytes.NewReader(loginBytes)
		r := httptest.NewRequest(http.MethodPost, "/login", loginReq)
		wr := httptest.NewRecorder()
		Env.Login(wr, r)
		updateUser := structs.User{FirstName: "Update", LastName: "Update", NickName: "Update", Email: repEmail, DateOfBirth: "06-08-2002", AboutMe: "Update", Avatar: "Update", IsPublic: 1}
		repBytes, err := json.Marshal(updateUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		repReq := bytes.NewReader(repBytes)
		request := httptest.NewRequest(http.MethodPut, "/updateprofileinfo", repReq)
		profW := httptest.NewRecorder()
		request.Header = http.Header{"Cookie": wr.Header()["Set-Cookie"]} // Now use this request for the validate handler
		Env.ProfileChange(profW, request)
		got := profW.Body.String()
		want := `{"Message":"Updated"}`
		if got != want {
			t.Errorf("Got %v Want %v", got, want)
		}
	})
	t.Run("Test with taken email", func(t *testing.T) {
		// register a user
		Env := handler.Env{Env: database}
		handlerEmail := "handlertest@" + uuid.NewV4().String() + ".com"
		repEmail := "new@" + uuid.NewV4().String() + ".com"
		originalUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: repEmail, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.InsertUser(originalUser, *database)

		// Create the struct that will be inserted
		sampleUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: handlerEmail, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

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
		loginUser := structs.User{Email: sampleUser.Email, Password: sampleUser.Password}
		// Marhsal the struct to a slice of bytes
		loginBytes, err := json.Marshal(loginUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		loginReq := bytes.NewReader(loginBytes)
		r := httptest.NewRequest(http.MethodPost, "/login", loginReq)
		wr := httptest.NewRecorder()
		Env.Login(wr, r)
		updateUser := structs.User{FirstName: "Update", LastName: "Update", NickName: "Update", Email: repEmail, DateOfBirth: "06-08-2002", AboutMe: "Update", Avatar: "Update", IsPublic: 1}
		repBytes, err := json.Marshal(updateUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		repReq := bytes.NewReader(repBytes)
		request := httptest.NewRequest(http.MethodPut, "/updateprofileinfo", repReq)
		profW := httptest.NewRecorder()
		request.Header = http.Header{"Cookie": wr.Header()["Set-Cookie"]} // Now use this request for the validate handler
		Env.ProfileChange(profW, request)
		got := profW.Body.String()
		want := `{"Message":"Email already in use"}`
		if got != want {
			t.Errorf("Got %v Want %v", got, want)
		}
	})
}
