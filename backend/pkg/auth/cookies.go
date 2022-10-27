package auth

import (
	"backend/pkg/structs"
	"fmt"
	"net/http"
	"time"
)

// CreateCookie creates a cookie for the specified responsewriter
func CreateCookie(w http.ResponseWriter, email string, DB *structs.DB) error{
	var user structs.User
	err := GetUser("email", email, &user, *DB)
	if err != nil {
		return err
	}
	cookieName := user.UserId + "&" + user.Email + "&" + user.SessionId
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   cookieName,
		Expires: time.Now().Add(24 * time.Hour),
	})
	fmt.Println("-!-!-!-!-!-!-!-!-!-!-!-!----  Cookie Created: " + time.Now().String())
	return nil
}

//RemoveCookie removes a cookie with a specific name
func RemoveCookie(w http.ResponseWriter) {
	c := &http.Cookie{Name: "session_token", MaxAge: 0, Expires: time.Now()}
	http.SetCookie(w, c)
}