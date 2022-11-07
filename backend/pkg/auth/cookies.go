package auth

import (
	"log"
	"net/http"
	"time"

	"backend/pkg/structs"
)

// CreateCookie creates a cookie for the specified responsewriter
func CreateCookie(w http.ResponseWriter, email string, DB *structs.DB) error {
	var user structs.User
	err := GetUser("email", email, &user, *DB)
	if err != nil {
		return err
	}
	cookieName := user.UserId + "&" + user.Email + "&" + user.SessionId
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    cookieName,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	return nil
}

// RemoveCookie removes a cookie with a specific name
func RemoveCookie(w http.ResponseWriter) {
	// c := &http.Cookie{Name: "session_token", Value: "", Expires: time.Now()}
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now()})
}

func ValidateCookie(c *http.Cookie, database *structs.DB, w http.ResponseWriter) bool {
	cookie, cErr := SliceCookie(c.Value)
	if cErr != nil {
		log.Println("Error slicing the cookie")
		return false
	}
	valid, seshErr := CheckSession(cookie[2], cookie[0], *database)
	if seshErr != nil {
		log.Println("Error searching for session")
		return false
	}
	if !valid {
		RemoveCookie(w)
		return false
	}
	return true
}
