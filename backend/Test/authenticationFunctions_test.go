package Test

import (
	"backend/pkg/functions"
	"backend/pkg/handler"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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

		//Create a struct to get the result
		var resultUser handler.User
		//Create a request
		req := httptest.NewRequest(http.MethodPost, "/", testReq)
		w := httptest.NewRecorder()
		errBody := functions.GetBody(&resultUser, w, req)

		if errBody != nil {
			t.Errorf("Error getting body from request: %v", errBody)
		}

		got := resultUser
		want := sampleUser

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}


	})
	t.Run("Getting bad body from the request", func(t *testing.T) {
		//Create a get request so that there is not body
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		var resultUser handler.User

		err := functions.GetBody(resultUser, w, req)
		//If the error is nil the function hasn't registered there is an invalid body
		if err == nil {
			t.Errorf("Error expected instead got %v", err)
		}
	})

}