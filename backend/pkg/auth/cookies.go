package auth

import (
	"backend/pkg/structs"
	"net/http"
	"time"
)

// CreateCookie creates a cookie for the specified responsewriter
func CreateCookie(w http.ResponseWriter, email string, DB *structs.DB, user structs.User) {
	cookieName := user.UserId + "&" + user.Email + "&" + user.SessionId
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   cookieName,
		Expires: time.Now().Add(24 * time.Hour),
	})
}

//RemoveCookie removes a cookie with a specific name
func RemoveCookie(w http.ResponseWriter, cookieName string) {
	c := &http.Cookie{Name: "session_token", MaxAge: 0, Expires: time.Now()}
	http.SetCookie(w, c)
}