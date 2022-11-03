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

func TestRemoveCookie(t *testing.T) {
	testEmail := "cookie@" + uuid.NewV4().String()

	Env := handler.Env{Env: database}
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
	var result structs.User
	auth.GetUser("email", testEmail, &result, *database)
	auth.CreateCookie(recorder, testEmail, database) // Create the cookie

	recorderDeleted := httptest.NewRecorder() // Drop a cookie on the recorder.

	auth.RemoveCookie(recorderDeleted) // Now try removing the cookie
	requestDeleted := &http.Request{Header: http.Header{"Cookie": recorderDeleted.Header()["Set-Cookie"]}}
	cookie, err := requestDeleted.Cookie("session_token") // Check if the cookie has been removed
	got := cookie.Value
	want := ""
	if got != want {
		t.Errorf("cookie found when should be deleted")
	}
}

func TestCreateCookie(t *testing.T) {
	testEmail := "cookie@" + uuid.NewV4().String()
	// // Create the database struct
	Env := handler.Env{Env: database}
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
	var result structs.User
	auth.GetUser("email", testEmail, &result, *database)
	auth.CreateCookie(recorder, testEmail, database)

	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}
	cookie, err := request.Cookie("session_token")
	if err != nil {
		t.Errorf("Error accessing cookie")
	}
	// Get the userId & sessionId
	want := result.UserId + "&" + result.Email + "&" + result.SessionId // The cookie should be equal to the userId + email + sessionId
	got := cookie.Value
	if want != got {
		t.Errorf("Got %v. Want %v.", got, want)
	}
}
