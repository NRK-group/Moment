package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestRegistration(t *testing.T) {
	t.Run("Request with valid URL", func(t *testing.T) {
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/registration", nil)
		w := httptest.NewRecorder()

		Env.Registration(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Request with Bad URL", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		Env := handler.Env{Env: DB}

		req := httptest.NewRequest(http.MethodGet, "/badUrl", nil)
		w := httptest.NewRecorder()

		Env.Registration(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Inserting registration data to database", func(t *testing.T) {
		Env := handler.Env{Env: database}

		// Create the struct that will be inserted
		sampleUser := structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: "handlertest@" + uuid.NewV4().String() + ".com", Password: "TestPass",
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

		// Now check if the data is added by querying the database manually and getting the specific user

		var resultUser structs.User
		auth.GetUser("email", sampleUser.Email, &resultUser, *Env.Env)
		resultUser.UserId = "-"
		resultUser.CreatedAt = ""

		sampleUser.Password = strconv.FormatBool(auth.CheckPasswordHash(sampleUser.Password, resultUser.Password))
		if err != nil {
			t.Errorf("Error hashing the password %v", err)
		}
		resultUser.Password = "true"
		auth.Capitalise(&sampleUser)

		want := sampleUser
		got := resultUser

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
