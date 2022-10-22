package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"backend/pkg/handler"
)

func TestGetBody(t *testing.T) {
	t.Run("Getting valid body from the request", func(t *testing.T) {
		// Create the struct that will be inserted
		sampleUser := handler.User{
			FirstName: "GetBodyTest", LastName: "GetBodyTest",
		}

		// Marhsal the struct to a slice of bytes
		sampleUserBytes, err := json.Marshal(sampleUser)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}

		// Create the bytes into a reader
		testReq := bytes.NewReader(sampleUserBytes)

		// Create a struct to get the result
		var resultUser handler.User
		// Create a request
		req := httptest.NewRequest(http.MethodPost, "/", testReq)
		w := httptest.NewRecorder()
		errBody := handler.GetBody(&resultUser, w, req)

		if errBody != nil {
			t.Errorf("Error getting body from request: %v", errBody)
		}

		got := resultUser
		want := sampleUser

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
	t.Run("Getting invalid body from the request", func(t *testing.T) {
		// Create a get request so that there is not body
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		var resultUser handler.User

		err := handler.GetBody(resultUser, w, req)
		// If the error is nil the function hasn't registered there is an invalid body
		if err == nil {
			t.Errorf("Error expected instead got %v", err)
		}
	})
}

var VALID_PASSWORDS = []string{
	"ValidPassword123",
	"ValidPas",
	"ValidPassword123",
}

var INVALID_PASSWORDS = []string{
	"invalidpassword1",
	"INVALIDPASSWORD1",
	"invalid",
	"Invalidbecauseitistoolong",
}

func TestValidPassword(t *testing.T){
	t.Run("Testing with valid passwords", func(t *testing.T) {
		for _, v := range VALID_PASSWORDS {
			got := handler.ValidPassword(v)
			want := true
			if got != want {
				t.Errorf("expected: %v, got %v ", want, got)
			}
		}

	})
	t.Run("Testing with invalid passwords", func(t *testing.T) {
		for _, v := range INVALID_PASSWORDS {
			got := handler.ValidPassword(v)
			want := false
			if got != want {
				t.Errorf("expected: %v, got %v ", want, got)
			}
		}

	})
}