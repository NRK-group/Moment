package Test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/handler"
)

func TestHealthCheckImageUploadHttpPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/imageUpload", nil)
	w := httptest.NewRecorder()

	Env := handler.Env{Env: database}
	Env.ImageUpload(w, req)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}
