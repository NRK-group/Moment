package Test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/handler"
)

func TestUpdateProfile(t *testing.T) {
	database := DatabaseSetup()
	t.Run("Request with valid URL", func(t *testing.T) {
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/updateprofile", nil)
		w := httptest.NewRecorder()

		Env.Update(w, req)
		want := 200
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Request with Bad URL", func(t *testing.T) {
		// Create the database struct
		Env := handler.Env{Env: database}

		req := httptest.NewRequest(http.MethodGet, "/badUrl", nil)
		w := httptest.NewRecorder()

		Env.Update(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}
