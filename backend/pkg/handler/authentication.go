package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId       string `json:"UserId"`
	SessionId    string `json:"SessionId"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	NickName     string `json:"NickName"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	DateOfBirth  string `json:"DateOfBirth"`
	AboutMe      string `json:"AboutMe"`
	Avatar       string `json:"Avatar"`
	CreatedAt    string `json:"CreatedAt"`
	IsLoggedIn   int    `json:"IsLoggedIn"`
	IsPublic     int    `json:"IsPublic"`
	NumFollowers int    `json:"NumFollowers"`
	NumFollowing int    `json:"NumFollowing"`
	NumPosts     int    `json:"NumPosts"`
}

// Login is a handler which checks if a the login credentials input by a user are valid
func (DB *DB) Login(w http.ResponseWriter, r *http.Request) {
	//Confirm the path is correct
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		//Check the credentials are correct
	}
}

// Registration is a handler where all registration functions are done
func (DB *DB) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if registration is correct
	if r.Method == "POST" {
		var newUser User
		// Get the body of the request
		GetBody(&newUser, w, r)
		// Insert the new user into the database
		err := DB.InsertUser(newUser)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// InsertUser is a method that inserts a new user into the database
func (DB DB) InsertUser(newUser User) error {
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

// GetBody marshalls the body of a request into a struct
func GetBody(b interface{}, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&b) // unmarshall the userdata
	if err != nil {
		fmt.Print(err)
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return err
}

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

//

func (DB *DB) GetUser(email, password string) (bool, string) {
	//Query the db to see if a user exsists with the inpit email
	rows, err := DB.DB.Query(`SELECT (email, password) From User WHERE email = ?`, email)
	if err != nil {
		fmt.Println("Error querying the db: ", err)
		return false, "Error querying the db"
	}
	counter := 0

	for rows.Next() {
		counter++
	}

	if counter == 0 {
		return false, "Account not found"
	}


	//If not return false with msg


	//Check if the password input is correct

	//If not return false with msg



	return false, ""
}