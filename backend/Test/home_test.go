package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
)

func TestHome(t *testing.T) {
	// initialize the database struct with a mock database
	database := &handler.DB{DB: sqlite.CreateDatabase("./social_network_test.db")}

	request := httptest.NewRequest(http.MethodGet, "/", nil) // create a request
	response := httptest.NewRecorder()                       // create a response recorder
	database.Home(response, request)                         // call the handler

	// check the response
	t.Run("Home handler response", func(t *testing.T) {
		got := response.Body.String() // get the response body
		want := "Documentation"       // set the expected response body

		if got != want {
			// if the response body is not the expected one, then fail the test
			t.Errorf("got %q, want %q", got, want)
		}
	})

	// check if the status code is the expected one
	t.Run("Home handler status code", func(t *testing.T) {
		got := response.Code  // get the status code
		want := http.StatusOK // set the expected status code

		if got != want {
			// if the status code is not the expected one, then fail the test
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
