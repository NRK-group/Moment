package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// CheckCredentials accepts the email and password a user inputs and checks whether the login credentials are valid.
// A boolean value and a string message are returned specifiying if the login was successful
func CheckCredentials(email, password string, DB *structs.DB) (bool, string) {
	rows, err := DB.DB.Query(`SELECT password FROM User WHERE email = ?`, email)
	if err != nil {
		log.Println("Error querying the db: ", err)
		return false, "Error querying the db"
	}
	var pass string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&pass)
	}
	if pass == "" {
		return false, "Account not found"
	}
	if CheckPasswordHash(password, pass) {
		return true, "Valid Login"
	}
	return false, "Incorrect Password"
}

// InsertUser is a method that inserts a new user into the database
func InsertUser(newUser structs.User, DB structs.DB) error {
	newUser.UserId = uuid.NewV4().String()                                                                 // Create a uuid for the user Id
	stmt, err := DB.DB.Prepare(`INSERT INTO User values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`) // Create the sql INSERT statement
	if err != nil {
		log.Println("Error preparing inserting user into the db: ", err)
		return err
	}
	var hashErr error // Hash the password and get the current time
	if !ValidPassword(newUser.Password) {
		return errors.New("Invalid Password")
	}
	newUser.Password, hashErr = HashPassword(newUser.Password)
	if hashErr != nil {
		log.Print("Error hashing password", hashErr)
		return hashErr
	}
	newUser.CreatedAt = time.Now().String()
	_, execErr := stmt.Exec(newUser.UserId, "-", newUser.FirstName, newUser.LastName, newUser.NickName, strings.ToLower(newUser.Email), newUser.DateOfBirth, newUser.Avatar, newUser.AboutMe, newUser.CreatedAt, newUser.IsLoggedIn, newUser.IsPublic, newUser.NumFollowers, newUser.NumFollowing, newUser.NumPosts, newUser.Password)
	if execErr != nil {
		log.Println("Error executing inserting user into the db: ", execErr)
		if strings.Contains(execErr.Error(), "UNIQUE constraint failed: User.email") {
			return errors.New("Email already in use")
		}
	}
	return execErr
}

// Delete is used to delet a row from a specefied table
func Delete(table, where, value string, DB structs.DB) error {
	dlt := "DELETE FROM " + table + " WHERE " + where
	stmt, err := DB.DB.Prepare(dlt + " = (?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(value)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSessionId
func UpdateSessionId(email, value string, DB structs.DB) error {
	var result structs.User
	err := GetUser("email", email, &result, DB)
	if err != nil {
		log.Println("Error getting the user profile details in updatesessionId func: ", err)
		return err
	}
	stmt, err := DB.DB.Prepare(`UPDATE User SET sessionId = ? WHERE email = ?`) // Update the session ID in the user table
	if err != nil {
		log.Println("Error preparing inserting user into the db: ", err)
		return err
	}
	_, err = stmt.Exec(value, email)
	if err != nil {
		log.Println("Error executing update sessionID")
		return err
	}
	err = Delete("UserSessions", "userId", result.UserId, DB)
	if err != nil {
		return err
	}
	if value != "-" {
		stmt, err := DB.DB.Prepare(`INSERT INTO UserSessions values (?, ?, ?)`) // Add the value to the db
		if err != nil {
			log.Println("Error Preparing Delete statement")
			return err
		}
		stmt.Exec(value, result.UserId, time.Now().String())
	}
	return nil
}

// Getuser is a function which queries the user table and gets the data from each column
func GetUser(datatype, value string, result *structs.User, DB structs.DB) error {
	rows, err := DB.DB.Query(`SELECT * FROM User WHERE `+datatype+` = ?`, value)
	if err != nil {
		log.Println("Error selecting data from db")
		return err
	}
	nothing := true
	defer rows.Close()
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
	if nothing { // No users were found
		return errors.New("No user found")
	}
	return nil
}

func CheckSession(session, user string, DB structs.DB) (bool, error) {
	rows, err := DB.DB.Query(`SELECT * FROM UserSessions WHERE sessionId = ? and userId = ?`, session, user)
	if err != nil {
		log.Println("Error selecting from the db")
		return false, err
	}
	valid := false
	defer rows.Close()
	for rows.Next() {
		valid = true
	}
	return valid, nil
}

// Update changes the specificed column in a specified table
func Update(table, set, to, where, id string, DB structs.DB) error {
	update := "UPDATE " + table + " SET " + set + " = '" + to + "' WHERE " + where + " = '" + id + "'"
	stmt, prepErr := DB.DB.Prepare(update)
	if prepErr != nil {
		log.Println("Error updating field: ", prepErr)
		return prepErr
	}
	_, err := stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserProfile changes the relevant row in the database when a user updates their profile
func UpdateUserProfile(userID string, user structs.User, DB structs.DB) error {
	qry, err := DB.DB.Prepare("UPDATE User SET firstName = ?, lastName = ?, nickName = ?, email = ?, DOB = ?, aboutMe = ?, avatar = ?, isPublic = ? WHERE userId = ?")
	if err != nil {
		log.Println("Error updating the database", err)
		return err
	}
	_, execErr := qry.Exec(user.FirstName, user.LastName, user.NickName, user.Email, user.DateOfBirth, user.AboutMe, user.Avatar, user.IsPublic, userID)
	if execErr != nil {
		log.Println("Error executing the update stmt ", execErr)
	}
	return execErr
}

// ActiveEmail checks if the email enetered is already in use and whether it belongs to the current user
func ActiveEmail(userID, email string, DB structs.DB) bool {
	if rows, err := DB.DB.Query("SELECT userId from User WHERE email = ?", email); err != nil {
		log.Println("Invalid Query")
		return true
	} else {
		var userId string
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userId)
			if userId != userID {
				return true
			}
		}
	}
	return false
}
