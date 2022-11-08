package Test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/handler"
)

func TestProfile(t *testing.T) {
	t.Run("Request with valid enpoint", func(t *testing.T) {
		// Create the database that will be used for testing

		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		w := httptest.NewRecorder()

		Env.Profile(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Request with invalid enpoint", func(t *testing.T) {
		// Create the database that will be used for testing
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
		w := httptest.NewRecorder()

		Env.Profile(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}
