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

func UpdateSessionId(email, value string, DB structs.DB) error {
	var result structs.User
	GetUser("email", email, &result, DB)
	stmt, err := DB.DB.Prepare(`UPDATE User SET sessionId = ? WHERE email = ?`) // Update the session ID in the user table
	if err != nil {
		fmt.Println("Error preparing inserting user into the db: ", err)
		return err
	}
	_, updateErr := stmt.Exec(value, email)
	if updateErr != nil {
		fmt.Println("Error executing update sessionID")
		return updateErr
	}
	if value == "-" {
		stmt, err := DB.DB.Prepare(`DELETE FROM UserSessions WHERE userId = ?`) // remove the session to the session table
		if err != nil {
			fmt.Println("Error Preparing Delete statement")
			return err
		}
		stmt.Exec(result.UserId)
	} else {
		stmt, err := DB.DB.Prepare(`INSERT INTO UserSessions values (?, ?, ?)`) // Add the value to the db
		if err != nil {
			fmt.Println("Error Preparing Delete statement")
			return err
		}
		stmt.Exec(value, result.UserId, time.Now().String())
	}
	return nil
}

func GetUser(datatype, value string, result *structs.User, DB structs.DB) error {
	rows, err := DB.DB.Query(`SELECT * FROM User WHERE `+datatype+` = ?`, value)
	if err != nil {
		fmt.Println("Error selecting data from db")
		return err
	}
	nothing := true
	for rows.Next() {
		nothing = false
		rows.Scan(
			&result.UserId,
			&result.SessionId,
			&result.FirstName,
			&result.LastName,
			&result.NickName,
			&result.Email,
			&result.DateOfBirth,
			&result.Avatar,
			&result.AboutMe,
			&result.CreatedAt,
			&result.IsLoggedIn,
			&result.IsPublic,
			&result.NumFollowers,
			&result.NumFollowing,
			&result.NumPosts,
			&result.Password)
	}
	if nothing {//No users were found
		return errors.New("No user found")
	}
	return nil
}