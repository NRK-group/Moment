package auth

import (
	"errors"
	"net/mail"
	"strings"

	"backend/pkg/structs"

	"golang.org/x/crypto/bcrypt"
)

// ValidPassword checks that the password input is valid and passes the requirements
func ValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 { // Check the length of the password is valid
		return false
	}
	if strings.ToLower(password) == password || strings.ToUpper(password) == password { // Check the password contains lower and uppercase values
		return false
	}
	return true
}

// HashPassword hashes a string so it cannot be read
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash Checks if a plaintext string and a hashed string are the same
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidEmail accepts an email (string) and checks if the syntax is correct for an email and returns a boolean
func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	// now check if there is no full stop
	split := strings.Split(email, "@")
	if len(split) != 2 {
		return false
	}
	return strings.Contains(split[1], ".")
}

// SliceCookie takes in the cookie and checks if it is valid to slice. If it can be a sliced a a slice of strings is returned containing the userId, email and sessionId
func SliceCookie(cookie string) ([]string, error) {
	if strings.Contains(cookie, "&") {
		emailSlc := strings.Split(cookie, "&")
		if len(emailSlc) == 3 {
			return emailSlc, nil
		}
	}
	return []string{}, errors.New("Invalid Cookie")
}

// ValidateValues checks that all input values are valid to be inserted
func ValidateValues(first, last, email, password string, isPublic int) (string, bool) {
	if len(first) == 0 || len(last) == 0 {
		return "Names must only contain letters and hyphens", false
	}
	if !ValidEmail(email) {
		return "Invalid email", false
	}
	if !ValidPassword(password) {
		return "Passwords must contain 8-16 lower and uppercase letters and special characters", false
	}
	if !(isPublic == 0 || isPublic == 1) {
		return "Invalid Account privacy", false
	} 
	return "", true
}

// Capitalise accepts a user struct and makes the neccassary fields capitalised
func Capitalise(obj *structs.User) {
	obj.FirstName = strings.Title(strings.ToLower(obj.FirstName))
	obj.LastName = strings.Title(strings.ToLower(obj.LastName))
	obj.NickName = strings.Title(strings.ToLower(obj.NickName))
	obj.Email = strings.ToLower(obj.Email)
}
