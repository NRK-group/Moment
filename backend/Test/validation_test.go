package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestGetBody(t *testing.T) {
	t.Run("Getting valid body from the request", func(t *testing.T) {
		// Create the struct that will be inserted
		sampleUser := structs.User{
			FirstName: "GetBodyTest", LastName: "GetBodyTest",
		}

		// Marhsal the struct to a slice of bytes
		sampleUserBytes, err := json.Marshal(sampleUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)

		// Create a struct to get the result
		var resultUser structs.User
		// Create a request
		req := httptest.NewRequest(http.MethodPost, "/", testReq)
		w := httptest.NewRecorder()
		errBody := handler.GetBody(&resultUser, w, req)

		if errBody != nil {
			t.Errorf("Error getting body from request: %v", errBody)
		}

		got := resultUser
		want := sampleUser

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
	t.Run("Getting invalid body from the request", func(t *testing.T) {
		// Create a get request so that there is not body
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		var resultUser structs.User

		err := handler.GetBody(resultUser, w, req)
		// If the error is nil the function hasn't registered there is an invalid body
		if err == nil {
			t.Errorf("Error expected instead got %v", err)
		}
	})
}

var VALID_PASSWORDS = []string{
	"ValidPassword123",
	"ValidPas",
	"ValidPassword123",
}

var INVALID_PASSWORDS = []string{
	"invalidpassword1",
	"INVALIDPASSWORD1",
	"invalid",
	"Invalidbecauseitistoolong",
}

func TestValidPassword(t *testing.T) {
	t.Run("Testing with valid passwords", func(t *testing.T) {
		for _, v := range VALID_PASSWORDS {
			got := auth.ValidPassword(v)
			want := true
			if got != want {
				t.Errorf("expected: %v, got %v ", want, got)
			}
		}
	})
	t.Run("Testing with invalid passwords", func(t *testing.T) {
		for _, v := range INVALID_PASSWORDS {
			got := auth.ValidPassword(v)
			want := false
			if got != want {
				t.Errorf("expected: %v, got %v ", want, got)
			}
		}
	})
}

func TestCreateCookie(t *testing.T) {
	testEmail = "cookie@" + uuid.NewV4().String()
	// Create a new user and log them in
	// Create the database that will be used for testing
	database := sqlite.CreateDatabase("./social_network_test.db")

	// migrate the database
	sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

	// Create the database struct
	DB := &structs.DB{DB: database}
	Env := handler.Env{Env: DB}
	inputUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: testEmail, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err := auth.InsertUser(*inputUser, *Env.Env)
	if err != nil {
		t.Errorf("Error inserting test struct")
	}

	// Create the struct that will be inserted
	sampleUser := &structs.User{
		Email: testEmail, Password: "Password123",
	}
	// Marshal the struct to bytes
	sampleUserBytes, err := json.Marshal(sampleUser)
	if err != nil {
		t.Errorf("Error marshalling the sampleUser")
	}

	// Create the bytes into a reader
	testReq := bytes.NewReader(sampleUserBytes)

	req := httptest.NewRequest(http.MethodPost, "/login", testReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)

	recorder := httptest.NewRecorder() // Drop a cookie on the recorder.
	auth.CreateCookie(recorder, testEmail, DB)

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	cookie, err := request.Cookie("session_token") 
	if err != nil {
		t.Errorf("Error accessing cookie")
	}

	// The cookie should be equal to the userId + email + sessionId

	// Get the userId & sessionId
	var result structs.User
	auth.GetUser("email", testEmail, &result, *DB)
	want := result.UserId + "&" + result.Email + "&" + result.SessionId
	got := cookie.Value
	if want != got {
		t.Errorf("Got %v. Want %v.", got, want)
	}
}
