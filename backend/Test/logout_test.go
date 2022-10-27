package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestLogout(t *testing.T) {
	t.Run("Test with a valid path", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}

		req := httptest.NewRequest(http.MethodGet, "/logout", nil)
		w := httptest.NewRecorder()

		Env.Logout(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Test with a invalid path", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}

		req := httptest.NewRequest(http.MethodGet, "/InvalidPath", nil)
		w := httptest.NewRecorder()

		Env.Logout(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Remove cookie", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}
		randEmail := uuid.NewV4().String()
		sampleUser := &structs.User{
			FirstName: "LogoutRemove", LastName: "LogoutRemove", NickName: "LogoutRemove", Email: randEmail, Password: "LogoutRemove1",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		fmt.Println("INSERTING FOR THIS TEST")
		err := auth.InsertUser(*sampleUser, *DB)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}

		// Now Log the user in
		// Create the struct that will be inserted
		testUser := &structs.User{
			Email: randEmail, Password: "logoutRemove1",
		}
		// Marshal the struct to bytes
		sampleUserBytes, err := json.Marshal(testUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)

		req := httptest.NewRequest(http.MethodPost, "/login", testReq)
		w := httptest.NewRecorder()
		Env.Login(w, req)
		// request the logout handler
		// request1 := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}

		// c, err := request1.Cookie("session_token")
		// if err != nil {
		// 	t.Errorf("Cookie Name: %v", err)
		// 	return
		// }
		// fmt.Println("Cookie VALUE : ", c.Value)

		// Env.Logout(w, req)

		request := &http.Request{Header: http.Header{"Cookie": w.Header()["Set-Cookie"]}}
		cookie, err := request.Cookie("session_token")
		if err != nil {
			t.Errorf("Error accessing the cookie: %v", err)
		}

		got := cookie.Value
		want := ""
		if got != want {
			t.Errorf("got: %v. Want: %v.", got, want)
		}
	})
}
