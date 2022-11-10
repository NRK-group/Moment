package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/handler"
	"backend/pkg/structs"
)

func TestGetBody(t *testing.T) {
	t.Run("Getting valid body from the request", func(t *testing.T) {
		// Create the struct that will be inserted
		sampleUser := structs.User{
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
		var resultUser structs.User
		// Create a request
		req := httptest.NewRequest(http.MethodPost, "/", testReq)
		w := httptest.NewRecorder()
		errBody := handler.GetBody(&resultUser, w, req)

		if errBody != nil {
			t.Errorf("Error getting body from request: %v", errBody)
		}

		got := resultUser
		want := sampleUser

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
	t.Run("Getting invalid body from the request", func(t *testing.T) {
		// Create a get request so that there is not body
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		var resultUser structs.User

		got := handler.GetBody(resultUser, w, req)
		var want error
		want = nil
		// If the error is nil the function hasn't registered there is an invalid body
		if got == want {
			t.Errorf("Error expected instead got %v", got)
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

func TestValidPassword(t *testing.T) {
	t.Run("Testing with valid passwords", func(t *testing.T) {
		for _, v := range VALID_PASSWORDS {
			got := auth.ValidPassword(v)
			want := true
			if got != want {
				t.Errorf("expected: %v, got %v ", want, got)
			}
		}
	})
	t.Run("Testing with invalid passwords", func(t *testing.T) {
		for _, v := range INVALID_PASSWORDS {
			got := auth.ValidPassword(v)
			want := false
			if got != want {
				t.Errorf("expected: %v, got %v ", want, got)
			}
		}
	})
}

var (
	emailtests    = []string{"valid@email.com", "valid@email.co.uk"}
	invalidEmails = []string{"invalidemail.c", "invalid@email", "Invalid.email@com", "Invalid"}
)

func TestValidEmail(t *testing.T) {
	t.Run("Testing valid emails", func(t *testing.T) {
		for _, v := range emailtests {
			got := auth.ValidEmail(v)
			want := true
			if got != want {
				t.Errorf("Error cehcking valid email: got %v -- want %v", got, want)
			}
		}
	})
	t.Run("Testing valid emails", func(t *testing.T) {
		for _, v := range invalidEmails {
			got := auth.ValidEmail(v)
			want := false
			if got != want {
				t.Errorf("Error cehcking valid email: got %v -- want %v ===== %v", got, want, v)
			}
		}
	})
}

var invalid = [][]string{{"", "Last", "email@email.com", "Password123"}, {"first", "", "email@email.com", "Password123"}, {"first", "Last", "email@ema", "Password123"}, {"first", "Last", "email@email.com", "assword123"}}

func TestValidValues(t *testing.T) {
	t.Run("Valid input values", func(t *testing.T) {
		_, got := auth.ValidateValues("Valid", "Valid", "email@email.com", "Password123", 0)
		want := true
		if got != want {
			t.Errorf("got %v, want %v ", got, want)
		}
	})
	t.Run("Invalid values for each", func(t *testing.T) {
		for _, v := range invalid {
			_, got := auth.ValidateValues(v[0], v[1], v[2], v[3], 1)
			want := false
			if got != want {
				t.Errorf("got %v, want %v ", got, want)
			}
		}
	})
}

func TestCapitalised(t *testing.T) {
	user := structs.User{FirstName: "first", LastName: "last", NickName: "nick", Email: "EmAil@eMAil.com"}
	auth.Capitalise(&user)

	got := user
	want := structs.User{FirstName: "First", LastName: "Last", NickName: "Nick", Email: "email@email.com"}
	if got != want {
		t.Errorf("got %v, want %v ", got, want)
	}
}