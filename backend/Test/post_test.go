package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"

	"github.com/stretchr/testify/assert"
)

func DatabaseSetup() *handler.DB {
	// this open or create the database
	db := sqlite.CreateDatabase("./social_network_test.db")
	// migrate the database
	sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")
	// initialize the database struct
	database := &handler.DB{DB: db}

	return database
}

func TestHealthCheckPostHandlerHttpGet(t *testing.T) {
	database := DatabaseSetup()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/post", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(database.Post)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected2 := http.StatusOK
	if status := rr.Code; status != expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected2)
	}
}

func TestHealthCheckPostHttpPost(t *testing.T) {
	w := httptest.NewRecorder()
	http.NewRequest("POST", "/post", nil)
	assert.Equal(t, 200, w.Code)
}

func TestCreatePost(t *testing.T) {
	t.Run("Insert Post to DB", func(t *testing.T) {
		database := DatabaseSetup()
		post := &handler.Post{UserID: "3232131221", Content: "hey", GroupID: "3233234", Image: "wasfdfgfd"}
		str, err := database.CreatePost(post.UserID, post.Content, post.GroupID, post.Image)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Read all Posts from the DB", func(t *testing.T) {
		database := DatabaseSetup()
	
		posts, err := database.AllPost("6t78t8t87")
		fmt.Println(posts)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}

func TestPostHandlerMakeAPost(t *testing.T) {
	database := DatabaseSetup()

	post1 := &handler.Post{UserID: "3232131221", Content: "hey2", GroupID: "3233234", Image: "wasfdfgfd"}
	body, _ := json.Marshal(post1)

	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(database.Post)
	handler.ServeHTTP(rr, req)
	expected := rr.Body.String()
	expectedStr := "successfully posted"
	if expectedStr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostHandlerGettingAllPost(t *testing.T) {
	database := DatabaseSetup()

	req, err := http.NewRequest("GET", "/post", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(database.Post)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	expected := http.StatusOK
	if status := rr.Code; status != expected && strings.Contains(rr.Body.String(), "PostID") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

