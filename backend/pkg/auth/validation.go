package auth

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// ValidPassword checks that the password input is valid and passes the requirements
func ValidPassword(password string) bool {
	// Check the length of the password is valid
	if len(password) < 8 || len(password) > 16 {
		return false
	}
	// Check the password contains lower and uppercase values
	if strings.ToLower(password) == password || strings.ToUpper(password) == password {
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
