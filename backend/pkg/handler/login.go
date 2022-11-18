package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"backend/pkg/auth"
	"backend/pkg/response"
	"backend/pkg/structs"

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
		successfulLogin, validationMsg := auth.CheckCredentials(strings.ToLower(userLogin.Email), userLogin.Password, DB.Env) // Validate the login creds
		if !successfulLogin {
			response.WriteMessage("Invalid credentials when logging in", validationMsg, w)
			return
		}
		sessionErr := auth.UpdateSessionId(strings.ToLower(userLogin.Email), uuid.NewV4().String(), *DB.Env) // Create a sessionID
		if sessionErr != nil {
			response.WriteMessage("Error creating the session", validationMsg, w)
			return
		}
		err = auth.CreateCookie(w, strings.ToLower(userLogin.Email), DB.Env) // Create the cookie
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		var result structs.User
		auth.GetUser("email", userLogin.Email, &result, *DB.Env)
		result.Message = validationMsg
		result.Password =""
		data, marshErr := json.Marshal(result)
		if marshErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
		}
		w.Write(data)
		return
	}
}
