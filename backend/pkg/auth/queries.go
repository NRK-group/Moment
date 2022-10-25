package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// CheckCredentials validates the credentials entered by the user
func CheckCredentials(email, password string, DB *structs.DB) (bool, string) {
	// Query the db to see if a user exsists with the inpit email
	rows, err := DB.DB.Query(`SELECT password FROM User WHERE email = ?`, email)
	if err != nil {
		fmt.Println("Error querying the db: ", err)
		return false, "Error querying the db"
	}
	counter := 0
	var pass string
	for rows.Next() {
		counter++
		rows.Scan(&pass)
	}
	// If not return false with msg
	if counter == 0 {
		return false, "Account not found"
	}
	fmt.Println("PASSWORD === ", pass)
	// Check if the password input is correct
	if CheckPasswordHash(password, pass) {
		return true, "Valid Login"
	}
	// If not return false with msg
	return false, "Incorrect Password"
}

// InsertUser is a method that inserts a new user into the database
func InsertUser(newUser structs.User, DB structs.DB) error {
	// Create a uuid for the user Id
	newUser.UserId = uuid.NewV4().String()
	// Create the sql INSERT statement
	stmt, err := DB.DB.Prepare(`INSERT INTO User values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println("Error preparing inserting user into the db: ", err)
		return err
	}
	// Hash the password and get the current time
	var hashErr error
	if !ValidPassword(newUser.Password) {
		return errors.New("Invalid Password")
	}
	newUser.Password, hashErr = HashPassword(newUser.Password)
	if hashErr != nil {
		log.Print("Error hashing password", hashErr)
		return hashErr
	}
	newUser.CreatedAt = time.Now().String()

	_, execErr := stmt.Exec(newUser.UserId, "-", newUser.FirstName, newUser.LastName, newUser.NickName, newUser.Email, newUser.DateOfBirth, newUser.Avatar, newUser.AboutMe, newUser.CreatedAt, newUser.IsLoggedIn, newUser.IsPublic, newUser.NumFollowers, newUser.NumFollowing, newUser.NumPosts, newUser.Password)
	if execErr != nil {
		fmt.Println("Error executing inserting user into the db: ", execErr)
		return execErr
	}
	return nil
}
