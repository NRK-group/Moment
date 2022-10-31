package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

var (
	logTestEmail   = "validCreds@" + uuid.NewV4().String()
	loginAttempts = [][]string{{"Account not found", "InvalidEmail@false.com", "Password123"}, {"Valid Login", logTestEmail, "Password123"}, {"Incorrect Password", logTestEmail, "IncorrectPassword"}}
)

func TestLogin(t *testing.T) {
	t.Run("Request with valid enpoint", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}

		req := httptest.NewRequest(http.MethodGet, "/login", nil)
		w := httptest.NewRecorder()

		Env.Login(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Request with invalid enpoint", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}

		req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
		w := httptest.NewRecorder()

		Env.Login(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Login with valid and invalid credentials", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: logTestEmail, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*inputUser, *Env.Env)
		if err != nil {
			t.Errorf("Error inserting test struct")
		}

		for _, value := range loginAttempts {

			// Create the struct that will be inserted
			sampleUser := &structs.User{
				Email: value[1], Password: value[2],
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
			want := value[0] 
			got := w.Body.String()
			
			if got != want {
				t.Errorf("got: %v. Want: %v.", got, want)
			}
		}
	})
	t.Run("Create a session ID for the logged in user", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}
		sampleUser := &structs.User{
			Email: logTestEmail, Password: "Password123",
		}
		sampleUserBytes, err := json.Marshal(sampleUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
		req := httptest.NewRequest(http.MethodPost, "/login", testReq)
		w := httptest.NewRecorder()
		Env.Login(w, req)
		//Check if sessionId for the user has been created
		rows, sessionErr := Env.Env.DB.Query(`SELECT sessionId FROM User WHERE email = ?`, logTestEmail)
		if sessionErr != nil {
			t.Errorf("Error selecting sessionId from the database")
		}
		var got string
		for rows.Next(){
			rows.Scan(&got)
		}
		notWant := "-"
		if got == notWant {
			t.Errorf("got %v", got)
		}
	})
}
