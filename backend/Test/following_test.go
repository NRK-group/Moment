package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/closefriend"
	"backend/pkg/follow"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestFollowing(t *testing.T) {
	t.Run("User is following the profile", func(t *testing.T) {
		// Register two users
		Env := handler.Env{Env: database}
		emailOne := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
		emailTwo := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
		// Create the struct that will be inserted
		firstUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailOne, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		// Marhsal the struct to a slice of bytes
		sampleUserBytes, err := json.Marshal(firstUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)
		req := httptest.NewRequest(http.MethodPost, "/registration", testReq)
		w := httptest.NewRecorder()
		Env.Registration(w, req)
		// Create the struct that will be inserted
		secondUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailTwo, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		// Marhsal the struct to a slice of bytes
		secondUserBytes, errSec := json.Marshal(secondUser)
		if errSec != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		// Create the bytes into a reader
		secReq := bytes.NewReader(secondUserBytes)
		secondReq := httptest.NewRequest(http.MethodPost, "/registration", secReq)
		sw := httptest.NewRecorder()
		Env.Registration(sw, secondReq)
		// Get the two users
		var userOne, userTwo structs.User
		auth.GetUser("email", emailOne, &userOne, *Env.Env)
		auth.GetUser("email", emailTwo, &userTwo, *Env.Env)
		follow.InsertFollow(userOne.UserId, userTwo.UserId, Env.Env)

		loginUser := structs.User{Email: emailOne, Password: "TestPass"}
		loginBytes, _ := json.Marshal(loginUser)
		loginRead := bytes.NewReader(loginBytes)

		loginW := httptest.NewRecorder()
		loginR := httptest.NewRequest(http.MethodPost, "/login", loginRead)
		Env.Login(loginW, loginR)

		// qryValues.Add("userID", userOne.UserId)
		checkReq := httptest.NewRequest(http.MethodGet, "/following", nil)
		qryValues := checkReq.URL.Query()
		qryValues.Add("followingID", userTwo.UserId)
		checkReq.URL.RawQuery = qryValues.Encode()
		checkReq.Header = http.Header{"Cookie": loginW.Header()["Set-Cookie"]}
		followW := httptest.NewRecorder()
		Env.Following(followW, checkReq)
		got := followW.Body.String()
		want := `{"Message":"Following"}`
		if got != want {
			t.Errorf("got %v \n want %v", got, want)
		}
	})
	t.Run("user1 not following user 2", func(t *testing.T) {
		// Register two users
		Env := handler.Env{Env: database}
		emailOne := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
		emailTwo := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")

		// Create the struct that will be inserted
		firstUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailOne, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		// Marhsal the struct to a slice of bytes
		sampleUserBytes, err := json.Marshal(firstUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)

		req := httptest.NewRequest(http.MethodPost, "/registration", testReq)
		w := httptest.NewRecorder()
		Env.Registration(w, req)
		// Create the struct that will be inserted
		secondUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailTwo, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		// Marhsal the struct to a slice of bytes
		secondUserBytes, errSec := json.Marshal(secondUser)
		if errSec != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		secReq := bytes.NewReader(secondUserBytes)

		secondReq := httptest.NewRequest(http.MethodPost, "/registration", secReq)
		sw := httptest.NewRecorder()
		Env.Registration(sw, secondReq)
		// Get the two users
		var userOne, userTwo structs.User
		auth.GetUser("email", emailOne, &userOne, *Env.Env)
		auth.GetUser("email", emailTwo, &userTwo, *Env.Env)
		// follow.InsertFollow(userOne.UserId, userTwo.UserId, Env.Env)

		loginUser := structs.User{Email: emailOne, Password: "TestPass"}
		loginBytes, _ := json.Marshal(loginUser)
		loginRead := bytes.NewReader(loginBytes)

		loginW := httptest.NewRecorder()
		loginR := httptest.NewRequest(http.MethodPost, "/login", loginRead)
		Env.Login(loginW, loginR)

		// qryValues.Add("userID", userOne.UserId)
		checkReq := httptest.NewRequest(http.MethodGet, "/following", nil)
		qryValues := checkReq.URL.Query()
		qryValues.Add("followingID", userTwo.UserId)
		checkReq.URL.RawQuery = qryValues.Encode()
		checkReq.Header = http.Header{"Cookie": loginW.Header()["Set-Cookie"]}
		followW := httptest.NewRecorder()
		Env.Following(followW, checkReq)
		got := followW.Body.String()
		want := `{"Message":"Not Following"}`
		if got != want {
			t.Errorf("got %v \n want %v", got, want)
		}
	})
	t.Run("Check when pending", func(t *testing.T) {
		Env := handler.Env{Env: database}
		emailOne := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
		emailTwo := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
		// Create the struct that will be inserted
		firstUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailOne, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		// Marhsal the struct to a slice of bytes
		sampleUserBytes, err := json.Marshal(firstUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)
		req := httptest.NewRequest(http.MethodPost, "/registration", testReq)
		w := httptest.NewRecorder()
		Env.Registration(w, req)
		// Create the struct that will be inserted
		secondUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailTwo, Password: "TestPass",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		// Marhsal the struct to a slice of bytes
		secondUserBytes, errSec := json.Marshal(secondUser)
		if errSec != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		// Create the bytes into a reader
		secReq := bytes.NewReader(secondUserBytes)
		secondReq := httptest.NewRequest(http.MethodPost, "/registration", secReq)
		sw := httptest.NewRecorder()
		Env.Registration(sw, secondReq)
		// Get the two users
		var userOne, userTwo structs.User
		auth.GetUser("email", emailOne, &userOne, *Env.Env)
		auth.GetUser("email", emailTwo, &userTwo, *Env.Env)
		follow.InsertFollowNotif(userOne.UserId, userTwo.UserId, "pending", Env.Env)

		loginUser := structs.User{Email: emailOne, Password: "TestPass"}
		loginBytes, _ := json.Marshal(loginUser)
		loginRead := bytes.NewReader(loginBytes)

		loginW := httptest.NewRecorder()
		loginR := httptest.NewRequest(http.MethodPost, "/login", loginRead)
		Env.Login(loginW, loginR)

		// qryValues.Add("userID", userOne.UserId)
		checkReq := httptest.NewRequest(http.MethodGet, "/following", nil)
		qryValues := checkReq.URL.Query()
		qryValues.Add("followingID", userTwo.UserId)
		checkReq.URL.RawQuery = qryValues.Encode()
		checkReq.Header = http.Header{"Cookie": loginW.Header()["Set-Cookie"]}
		followW := httptest.NewRecorder()
		Env.Following(followW, checkReq)
		got := followW.Body.String()
		want := `{"Message":"Pending"}`
		if got != want {
			t.Errorf("got %v \n want %v", got, want)
		}
	})
	t.Run("Close Friends rel", func(t *testing.T) {
		Env := handler.Env{Env: database}
		userOne := CreateUser(database, t)
		userTwo := CreateUser(database, t)
		// Login to userOne
		loginUser := structs.User{Email: userOne.Email, Password: "Password123"}
		sampleUserBytes, _ := json.Marshal(loginUser)
		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)
		r := httptest.NewRequest(http.MethodPost, "/login", testReq)
		w := httptest.NewRecorder()
		Env.Login(w, r)
		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/following", nil)
		req.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
		values := req.URL.Query()
		values.Add("followingID", userTwo.UserId)
		req.URL.RawQuery = values.Encode()
		follow.InsertFollow(userOne.UserId, userTwo.UserId, Env.Env)
		closefriend.AddCloseFriend(userTwo.UserId, userOne.UserId, *Env.Env)
		Env.Following(wr, req)
		got := wr.Body.String()
		want := `{"Message":"Close Friend"}`
		if got != want {
			t.Errorf("Got = %v Want = %v", got, want)
		}
	})
}
