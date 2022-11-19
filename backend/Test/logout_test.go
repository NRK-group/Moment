package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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
		// database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		// DB := &structs.DB{DB: database}
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodPut, "/logout", nil)
		w := httptest.NewRecorder()

		Env.Logout(w, req)
		want := 400
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

		req := httptest.NewRequest(http.MethodPut, "/InvalidPath", nil)
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
		// database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		// DB := &structs.DB{DB: database}
		Env := handler.Env{Env: database}
		randEmail := "removeCookie@" + uuid.NewV4().String() + ".com"
		sampleUser := &structs.User{
			FirstName: "LogoutRemove", LastName: "LogoutRemove", NickName: "LogoutRemove", Email: randEmail, Password: "LogoutRemove1",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}

		// Now Log the user in
		// Create the struct that will be inserted
		testUser := &structs.User{
			Email: randEmail, Password: "LogoutRemove1",
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
		auth.CreateCookie(w, testUser.Email, database)
		// request the logout handler
		req.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}

		_, err = req.Cookie("session_token")
		if err != nil {
			t.Errorf("Cookie Name: %v", err)
			return
		}

		request := httptest.NewRequest(http.MethodGet, "/logout", nil)
		request.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
		wr := httptest.NewRecorder()
		Env.Logout(wr, request)

		result := &http.Request{Header: http.Header{"Cookie": wr.Header()["Set-Cookie"]}}
		cookie, err := result.Cookie("session_token")
		if err != nil {
			t.Errorf("Error accessing the cookie: %v", err)
			return
		}

		got := cookie.Value
		want := ""
		if got != want {
			t.Errorf("got: %v. Want: %v.", got, want)
		}
	})
	t.Run("Remove user from session table", func(t *testing.T) {
		// Create the database that will be used for testing
		Env := handler.Env{Env: database}
		randEmail := "removeCookie@" + uuid.NewV4().String() + ".com"
		sampleUser := &structs.User{
			FirstName: "LogoutRemove", LastName: "LogoutRemove", NickName: "LogoutRemove", Email: strings.ToLower(randEmail), Password: "LogoutRemove1",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}

		// Now Log the user in
		// Create the struct that will be inserted
		testUser := &structs.User{
			Email: randEmail, Password: "LogoutRemove1",
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
		// now check if the session is in the table
		err = auth.GetUser("email", sampleUser.Email, sampleUser, *database)
		if err != nil {
			t.Errorf("Error getting user: %v", err)
		}
		rows, err := database.DB.Query("SELECT * FROM UserSessions WHERE userId = ?", sampleUser.UserId)
		if err != nil {
			t.Errorf("Error querying db:  %v", err)
		}
		before := 0
		for rows.Next() {
			before++
		}

		// Now log the user out
		// sampleUserBytes, err = json.Marshal(sampleUser)
		// if err != nil {
		// 	t.Errorf("Error marshalling the sampleUser")
		// }

		// Create the bytes into a reader
		// testReq = bytes.NewReader(sampleUserBytes)

		r := httptest.NewRequest(http.MethodGet, "/logout", nil)
		r.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}

		wr := httptest.NewRecorder()
		Env.Logout(wr, r)

		rows, err = database.DB.Query("SELECT * FROM UserSessions WHERE userId = ?", sampleUser.UserId)
		if err != nil {
			t.Errorf("Error querying db:  %v", err)
		}
		after := 0
		for rows.Next() {
			after++
		}
		if after == before {
			t.Errorf("Got: %v, Want: %v", after, before)
		}
	})
}
