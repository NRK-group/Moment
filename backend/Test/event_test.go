package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/pkg/group"
	"backend/pkg/handler"
	"backend/pkg/structs"
)



func TestHealthCheckEventHttpGet(t *testing.T) {
	database := DatabaseSetup()
	req := httptest.NewRequest(http.MethodGet, "/event", nil)
	w := httptest.NewRecorder()

	Env := handler.Env{Env: database}
	Env.Event(w, req)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}


func TestHealthCheckEventHttpPost(t *testing.T) {
	database := DatabaseSetup()
	req := httptest.NewRequest(http.MethodPost, "/event", nil)
	w := httptest.NewRecorder()

	Env := handler.Env{Env: database}
	Env.Event(w, req)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}


func TestCreateEvent(t *testing.T) {
	t.Run("Creating a Event", func(t *testing.T) {
		database := DatabaseSetup()
		event1 := structs.Event{Name: "Pie", Description: "Eating Pie", Admin: "wasfdfgfd"}
		str, err := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Read all Posts from the DB", func(t *testing.T) {
		database := DatabaseSetup()
		groups, err := group.AllGroups("6t78t8t87", database)
		fmt.Println(groups)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}


func TestHealthCheckEventHandlerHttpGet(t *testing.T) {
	database := DatabaseSetup()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/event", nil)
	if err != nil {
		t.Fatal(err)
	}
	Env := &handler.Env{Env: database}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Group)
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

func TestPostHandlerCreatingAEvent(t *testing.T) {
	database := DatabaseSetup()

	group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: "wasfdfgfd"}
	body, _ := json.Marshal(group1)

	req, err := http.NewRequest("POST", "/group", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	Env := &handler.Env{Env: database}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Group)
	handler.ServeHTTP(rr, req)
	expected := rr.Body.String()
	expectedStr := "successfully posted"
	if expectedStr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}



func TestPostHandlerGettingAllEventByGroup(t *testing.T) {
	database := DatabaseSetup()

	req, err := http.NewRequest("GET", "/event", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Env := &handler.Env{Env: database}
	handler := http.HandlerFunc(Env.Group)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	expected := http.StatusOK
	if status := rr.Code; status != expected && strings.Contains(rr.Body.String(), "GroupID") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
