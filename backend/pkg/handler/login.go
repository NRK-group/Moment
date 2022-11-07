package handler

import (
	"backend/pkg/auth"
	"backend/pkg/structs"
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
)
// Login is a handler that validates the credentials input by a user

func (DB *Env) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "POST" {
		var userLogin structs.User
		err := GetBody(&userLogin, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		successfulLogin, validationMsg := auth.CheckCredentials(userLogin.Email, userLogin.Password, DB.Env) // Validate the login creds
		if !successfulLogin {
			io.WriteString(w, validationMsg)
			return
		}
		sessionErr := auth.UpdateSessionId(userLogin.Email, uuid.NewV4().String(), *DB.Env) // Create a sessionID
		if sessionErr != nil {
			io.WriteString(w, validationMsg)
			return
		}
		err = auth.CreateCookie(w, userLogin.Email, DB.Env) // Create the cookie
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/text")
		io.WriteString(w, validationMsg)
		return
	}
}