package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"backend/pkg/auth"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func SetupCorsResponse(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

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

// Logout is a handler that runs all functions to logout the user
func (DB *Env) Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		fmt.Println("-------------------------!!!!____!!!!____!11 === LOGGING OUT 0 ")
		c, err := r.Cookie("session_token") // Access the cookie
		if err == nil {                     // Cookie is present so remove
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now()})
		} else { // The user isnt logged in
			http.Error(w, "401 unauthorized", http.StatusUnauthorized)
			return
		}
		emailSlc, slcErr := auth.SliceCookie(c.Value)
		if slcErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Update the sessionId update in users table and remove from sessions table
		err = auth.UpdateSessionId(emailSlc[1], "-", *DB.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		return
	}
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
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
