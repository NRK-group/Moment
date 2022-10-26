package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"backend/pkg/structs"
	"backend/pkg/post"
	"backend/pkg/handler"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckCommetsHandlerHttpGet(t *testing.T) {
	database := DatabaseSetup()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/post", nil)
	if err != nil {
		t.Fatal(err)
	}
	Env := &handler.Env{Env: database}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Post)
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

func TestHealthCheckCommetsHttpPost(t *testing.T) {
	w := httptest.NewRecorder()
	http.NewRequest("POST", "/post", nil)
	assert.Equal(t, 200, w.Code)
}

func TestCreateCommets(t *testing.T) {
	t.Run("Insert Post to DB", func(t *testing.T) {
		database := DatabaseSetup()
		post1 := structs.Post{UserID: "3232131221", Content: "hey", GroupID: "3233234", Image: "wasfdfgfd"}
		str, err := post.CreatePost(post1.UserID, post1.Content, post1.GroupID, post1.Image, database)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Read all Posts from the DB", func(t *testing.T) {
		database := DatabaseSetup()

		posts, err := post.AllPost("6t78t8t87", database)
		fmt.Println(posts)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}

func TestCommetstHandlerMakeACommet(t *testing.T) {
	database := DatabaseSetup()

	post1 := structs.Post{UserID: "3232131221", Content: "hey2", GroupID: "3233234", Image: "wasfdfgfd"}
	body, _ := json.Marshal(post1)

	req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	Env := &handler.Env{Env: database}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Post)
	handler.ServeHTTP(rr, req)
	expected := rr.Body.String()
	expectedStr := "successfully posted"
	if expectedStr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCommetsHandlerGettingAllCommet(t *testing.T) {
	database := DatabaseSetup()

	req, err := http.NewRequest("GET", "/post", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Env := &handler.Env{Env: database}
	handler := http.HandlerFunc(Env.Post)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	expected := http.StatusOK
	if status := rr.Code; status != expected && strings.Contains(rr.Body.String(), "PostID") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
