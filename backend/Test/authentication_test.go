package Test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
)

func TestRegistration(t *testing.T) {
	t.Run("Request with valid URL", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &handler.DB{DB: database}

		req := httptest.NewRequest(http.MethodPost, "/registration", nil)
		w := httptest.NewRecorder()

		DB.Registration(w, req)
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
		DB := &handler.DB{DB: database}

		req := httptest.NewRequest(http.MethodPost, "/badUrl", nil)
		w := httptest.NewRecorder()

		DB.Registration(w, req)
		want := 404
		got := w.Code

		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}
