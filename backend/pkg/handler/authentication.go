package handler

import (
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// Login is a handler that vlidates the credentials input by a user
func (DB *Env) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method == "POST" {
		var userLogin structs.User
		err := GetBody(&userLogin, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		successfulLogin, validationMsg := auth.CheckCredentials(userLogin.Email, userLogin.Password, DB.Env) // Validate the login creds
		if !successfulLogin {                                                                                // If credentials are invalid retrun unauthorized error and message
			http.Error(w, "401 Unauthorized ", http.StatusUnauthorized)
			return
		}
		sessionErr := auth.UpdateSessionId(userLogin.Email, uuid.NewV4().String(), *DB.Env) // Create a sessionID
		if sessionErr != nil {
			http.Error(w, "401 Unauthorized ", http.StatusUnauthorized)
			return
		}
		err = auth.CreateCookie(w, userLogin.Email, DB.Env)//Create the cookie
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(validationMsg))
		return
	}
}

// Registration is a handler where all registration functions are done
func (DB *Env) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Check if registration is correct
	if r.Method == "POST" {
		var newUser structs.User
		// Get the body of the request
		bodyErr := GetBody(&newUser, w, r)
		if bodyErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Insert the new user into the database
		err := auth.InsertUser(newUser, *DB.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
