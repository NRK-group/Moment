package auth

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// ValidPassword checks that the password input is valid and passes the requirements
func ValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 {// Check the length of the password is valid
		return false
	}
	if strings.ToLower(password) == password || strings.ToUpper(password) == password {// Check the password contains lower and uppercase values
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

//SliceCookie takes in the cookie and checks if it is valid to slice. If it can be a sliced a a slice of strings is returned containing the userId, email and sessionId
func SliceCookie(cookie string) ([]string, error) {
	if strings.Contains(cookie, "&") {
		emailSlc := strings.Split(cookie, "&")
		if len(emailSlc) == 3 {
			return emailSlc, nil
		}
	}
	return []string{}, errors.New("Invalid Cookie")
}
