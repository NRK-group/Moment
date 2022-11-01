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

func TestValidate(t *testing.T) {
	t.Run("Request with valid enpoint", func(t *testing.T) {
		database := DatabaseSetup()
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodPut, "/validate", nil)
		w := httptest.NewRecorder()

		Env.Validate(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Request with invalid enpoint", func(t *testing.T) {
		database := DatabaseSetup()
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodPut, "/invalid", nil)
		w := httptest.NewRecorder()

		Env.Validate(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Send a request from a valid user", func(t *testing.T) {
		database := DatabaseSetup()
		Env := handler.Env{Env: database}
		
		// Now the user is logged in Check if there is a valid session
		request := httptest.NewRequest(http.MethodGet, "/validate", nil)
		wr := httptest.NewRecorder()

		Env.Validate(wr, request)

		want := "401 Unauthorized\n"
		got := wr.Body.String()
		if got != want {
			t.Errorf("Got %v, want %v,", got, want)
		}
	})
	t.Run("Send a request from a valid user", func(t *testing.T) {
		validationValidEmail := "validCreds@" + uuid.NewV4().String()
		database := DatabaseSetup()
		Env := handler.Env{Env: database}
		inputUser := &structs.User{
			FirstName: "Validation", LastName: "Validation", NickName: "NickTest", Email: validationValidEmail, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*inputUser, *Env.Env)
		if err != nil {
			t.Errorf("Error inserting test struct")
		}

		sampleUser := &structs.User{
			Email: validationValidEmail, Password: "Password123",
		}
		sampleUserBytes, err := json.Marshal(sampleUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
		req := httptest.NewRequest(http.MethodPost, "/login", testReq)
		w := httptest.NewRecorder()
		Env.Login(w, req)
		// Now the user is logged in Cehck if there is a valid session

		request := httptest.NewRequest(http.MethodGet, "/validate", nil)
		request.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]} // Now use this request for the validate handler
		wr := httptest.NewRecorder()

		Env.Validate(wr, request)

		want := "Validated"
		got := wr.Body.String()
		if got != want {
			t.Errorf("Got %v, want %v,", got, want)
		}
	})
}
