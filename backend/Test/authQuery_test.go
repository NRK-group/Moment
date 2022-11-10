package Test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"backend/pkg/auth"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

var tests = []structs.User{
	{
		FirstName: "", LastName: "Length", NickName: "Length", Email: "Length" + uuid.NewV4().String(), Password: "Length",
		DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	},

	{
		FirstName: "Length", LastName: "", NickName: "Length", Email: "Length" + uuid.NewV4().String(), Password: "Length",
		DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	},

	{
		FirstName: "Length", LastName: "Length", NickName: "Length", Email: "", Password: "Length",
		DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	},

	{
		FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length" + uuid.NewV4().String(), Password: "",
		DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	},

	{
		FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length" + uuid.NewV4().String(), Password: "Length",
		DateOfBirth: "", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	},

	{
		FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length" + uuid.NewV4().String(), Password: "Length",
		DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	},
}

func TestInsertUser(t *testing.T) {
	randEmail := "insertUSer@" + uuid.NewV4().String()
	t.Run("Insert valid user to DB", func(t *testing.T) {
		sampleUser := structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: strings.ToLower(randEmail), Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(sampleUser, *database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		// Query the db to check if the user was inserted
		rows, err := database.DB.Query(`SELECT * FROM User WHERE Email = ?`, sampleUser.Email)
		// var userId, sessionId, firstName, lastName, nickName, email, DOB, avatar, aboutMe, createdAt, isLoggedIn, isPublic, numFollowers, numFollowing, numPosts, password string
		var resultUser structs.User

		for rows.Next() {
			rows.Scan(&resultUser.UserId, &resultUser.SessionId, &resultUser.FirstName, &resultUser.LastName, &resultUser.NickName, &resultUser.Email, &resultUser.DateOfBirth, &resultUser.Avatar, &resultUser.AboutMe, &resultUser.CreatedAt, &resultUser.IsLoggedIn, &resultUser.IsPublic, &resultUser.NumFollowers, &resultUser.NumFollowing, &resultUser.NumPosts, &resultUser.Password)
		}
		resultUser.UserId = "-"
		resultUser.CreatedAt = "-"
		sampleUser.Password = strconv.FormatBool(auth.CheckPasswordHash(sampleUser.Password, resultUser.Password))
		if err != nil {
			t.Errorf("Error hashing the password %v", err)
		}
		resultUser.Password = "true"
		want := sampleUser
		got := resultUser

		if got != want {
			t.Errorf("want %v, \n got %v", want, got)
		}
	})
	t.Run("inserting a user with used email to the db", func(t *testing.T) {
		sampleUser := &structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err == nil {
			t.Errorf("Error Catching already used email %v", err)
		}
	})
	t.Run("Check the length of neccesary values can't be 0", func(t *testing.T) {
		for _, s := range tests {
			err := auth.InsertUser(s, *database)
			if err == nil {
				t.Errorf("Error Catching empty values %v", err)
			}
		}
	})
}

func TestCheckCredentials(t *testing.T) {
	testEmail := "GetUser@" + uuid.NewV4().String()

	t.Run("Non-existing account entered", func(t *testing.T) {
		sampleUser := &structs.User{
			FirstName: "GetUser", LastName: "GetUser", NickName: "GetUser", Email: "123", Password: "GetUser",
			DateOfBirth: "GetUser", AboutMe: "GetUser", Avatar: "GetUser", CreatedAt: "GetUser", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		// Now check if the value is input to the db
		wantStr := "Account not found"
		wantBool := false

		gotBool, gotStr := auth.CheckCredentials(sampleUser.Email, sampleUser.Password, database)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}
	})

	t.Run("Check with correct credentials", func(t *testing.T) {
		sampleUser := &structs.User{
			FirstName: "GetUser", LastName: "GetUser", NickName: "GetUser", Email: testEmail, Password: "GetUser123",
			DateOfBirth: "GetUser", AboutMe: "GetUser", Avatar: "GetUser", CreatedAt: "GetUser", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error Inserting the test user to the db")
		}

		// Now check if the value is input to the db
		wantStr := "Valid Login"
		wantBool := true

		gotBool, gotStr := auth.CheckCredentials(strings.ToLower(sampleUser.Email), sampleUser.Password, database)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}
	})

	t.Run("Check with incorrect password", func(t *testing.T) {
		// Now check if the value is input to the db
		wantStr := "Incorrect Password"
		wantBool := false

		gotBool, gotStr := auth.CheckCredentials(testEmail, "incorrectPassword", database)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}
	})
}

func TestGetUser(t *testing.T) {
	randEmail := uuid.NewV4().String()
	t.Run("Getting valid user", func(t *testing.T) {
		// Create the database struct

		sampleUser := &structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		var got structs.User
		getErr := auth.GetUser(`email`, randEmail, &got, *database)
		if getErr != nil {
			t.Errorf("Error getting the user from the database")
		}
		got.CreatedAt = "-"
		got.UserId = "-"
		sampleUser.Password = strconv.FormatBool(auth.CheckPasswordHash(sampleUser.Password, got.Password))
		if err != nil {
			t.Errorf("Error hashing the password %v", err)
		}
		got.Password = "true"

		if got != *sampleUser {
			t.Errorf("got: %v. Want: %v.", got, *sampleUser)
		}
	})
	t.Run("Getting User that doesnt exsist", func(t *testing.T) {
		// Create the database struct
		var got structs.User
		getErr := auth.GetUser("email", "", &got, *database)
		if getErr == nil {
			t.Errorf("Error recognising invalid user details")
		}
	})
}

func TestUpdateSessionId(t *testing.T) {
	t.Run("Adding session to user in user table", func(t *testing.T) {
		// Create the database struct
		randEmail := uuid.NewV4().String()
		sampleUser := &structs.User{
			FirstName: "SessionTest", LastName: "SessionTest", NickName: "SessionTest", Email: randEmail, Password: "SessionTest",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}
		newSess := uuid.NewV4().String()
		auth.UpdateSessionId(randEmail, newSess, *database)
		var result structs.User
		getErr := auth.GetUser("email", randEmail, &result, *database)
		if getErr != nil {
			t.Errorf("Error getting the new user")
		}
		got := result.SessionId
		want := newSess

		if got != want {
			t.Errorf("Got: %v. Want: %v.", got, want)
		}
	})
	t.Run("adding the session to the session table", func(t *testing.T) {
		randEmail := uuid.NewV4().String()
		sampleUser := &structs.User{
			FirstName: "SessionTest", LastName: "SessionTest", NickName: "SessionTest", Email: randEmail, Password: "SessionTest",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}
		newSess := uuid.NewV4().String()
		auth.UpdateSessionId(randEmail, newSess, *database)
		var result structs.User
		getErr := auth.GetUser("email", randEmail, &result, *database)
		if getErr != nil {
			t.Errorf("Error getting the new user")
		}
		// Get the results from the session Id table
		rows, querErr := database.DB.Query(`SELECT * FROM UserSessions WHERE userId = ?`, result.UserId)
		if querErr != nil {
			t.Errorf("Error accessing the table")
			return
		}
		var gotsess, gotuser, date string
		for rows.Next() {
			rows.Scan(&gotsess, &gotuser, &date)
		}

		if gotsess != newSess || gotuser != result.UserId {
			t.Errorf("Got: %v %v. Want: %v, %v.", gotsess, gotuser, newSess, result.UserId)
		}
	})
	t.Run("Removing session from the user table", func(t *testing.T) {
		// Create the database struct
		randEmail := uuid.NewV4().String()
		sampleUser := &structs.User{
			FirstName: "SessionTest", LastName: "SessionTest", NickName: "SessionTest", Email: randEmail, Password: "SessionTest",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}
		auth.UpdateSessionId(randEmail, "-", *database)

		var result structs.User
		getErr := auth.GetUser("email", randEmail, &result, *database)
		if getErr != nil {
			t.Errorf("Error getting the new user")
		}
		got := result.SessionId
		want := "-"

		if got != want {
			t.Errorf("Got: %v. Want: %v.", got, want)
		}
	})

	t.Run("Removing the session from the session table", func(t *testing.T) {
		// Create the database struct
		randEmail := uuid.NewV4().String() // Create a new email
		sampleUser := &structs.User{
			FirstName: "SessionTest", LastName: "SessionTest", NickName: "SessionTest", Email: randEmail, Password: "SessionTest",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		} // Create a sample user to insert
		err := auth.InsertUser(*sampleUser, *database)
		if err != nil {
			t.Errorf("Error inserting the new user to the db")
		}
		newSess := uuid.NewV4().String()
		auth.UpdateSessionId(randEmail, newSess, *database) // Add the session to the db
		auth.UpdateSessionId(randEmail, "-", *database)     // Now get rid of the session from the session table
		var result structs.User
		auth.GetUser("email", randEmail, &result, *database)
		rows, _ := database.DB.Query(`SELECT * FROM UserSessions WHERE userId = ?`, result.UserId) // Check the row doesnt exsist
		counter := 0
		for rows.Next() {
			counter++
		}
		got := counter
		want := 0
		if counter != want {
			t.Errorf("Got: %v. Want: %v.", got, want)
		}
	})
}

func TestUpdate(t *testing.T) {
	randEmail := uuid.NewV4().String() // Create a new email
	sampleUser := &structs.User{
		FirstName: "SessionTest", LastName: "SessionTest", NickName: "SessionTest", Email: randEmail, Password: "SessionTest",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	} // Create a sample user to insert
	err := auth.InsertUser(*sampleUser, *database)
	if err != nil {
		t.Errorf("Error inserting the new user to the db")
	}

	testUID := uuid.NewV4().String()
	updateErr := auth.Update("User", "userId", testUID, "email", randEmail, *database)
	if updateErr != nil {
		t.Errorf("Error updating the table %v", updateErr)
	}
	var result structs.User
	err = auth.GetUser("email", randEmail, &result, *database)
	got := result.UserId
	want := testUID
	if got != want {
		t.Errorf("Got: %v. Want: %v.", got, want)
	}
}

var newEmail = "email@" + uuid.NewV4().String() + ".com"

var updateExamples = []structs.User{
	{FirstName: "", LastName: "Update", NickName: "Update", Email: strings.ToLower(newEmail), DateOfBirth: "06-08-2002", AboutMe: "Update", Avatar: "Update", IsPublic: 1},
	{FirstName: "Update", LastName: "", NickName: "Update", Email: strings.ToLower(newEmail), DateOfBirth: "06-08-2002", AboutMe: "Update", Avatar: "Update", IsPublic: 1},
	{FirstName: "Update", LastName: "Update", NickName: "Update", Email: "", DateOfBirth: "06-08-2002", AboutMe: "Update", Avatar: "Update", IsPublic: 1},
}

func TestUpdateUserProfile(t *testing.T) {
	t.Run("Valid update", func(t *testing.T) {
		updateEmail := "email@" + uuid.NewV4().String() + ".com"
		newEmail := "email@" + uuid.NewV4().String() + ".com"
		currTime := time.Now().String()

		firstUser := structs.User{FirstName: "First", LastName: "Last", NickName: "Nick", Email: strings.ToLower(updateEmail), Password: "Password123", DateOfBirth: currTime, AboutMe: "AboutMe", Avatar: "Test", IsPublic: 0}
		err := auth.InsertUser(firstUser, *database)
		if err != nil {
			t.Errorf("Error inserting the user")
			return
		}
		var temp structs.User
		getErr := auth.GetUser("email", strings.ToLower(updateEmail), &temp, *database)
		if getErr != nil {
			t.Errorf("Error getting the user")
			return
		}
		// Create the struct to update the user
		result := structs.User{FirstName: "Update", LastName: "Update", NickName: "Update", Email: strings.ToLower(newEmail), DateOfBirth: "06-08-2002", AboutMe: "Update", Avatar: "Update", IsPublic: 1}
		updateErr := auth.UpdateUserProfile(temp.UserId, result, *database)
		if updateErr != nil {
			t.Errorf("Error updating the user profile ")
			return
		}
		// Get the user to see if results have been updated
		var got structs.User
		auth.GetUser("email", strings.ToLower(newEmail), &got, *database)
		got.UserId = ""
		got.SessionId = ""
		got.Password = ""
		got.DateOfBirth = ""
		got.CreatedAt = ""
		result.DateOfBirth = ""

		if got != result {
			t.Errorf("Got %v want %v", got, result)
		}
	})
	t.Run("insert invalid values", func(t *testing.T) {
		updateEmail := "email@" + uuid.NewV4().String() + ".com"
		currTime := time.Now().String()

		firstUser := structs.User{FirstName: "First", LastName: "Last", NickName: "Nick", Email: strings.ToLower(updateEmail), Password: "Password123", DateOfBirth: currTime, AboutMe: "AboutMe", Avatar: "Test", IsPublic: 0}
		err := auth.InsertUser(firstUser, *database)
		if err != nil {
			t.Errorf("Error inserting the user")
			return
		}
		var temp structs.User
		getErr := auth.GetUser("email", strings.ToLower(updateEmail), &temp, *database)
		if getErr != nil {
			t.Errorf("Error getting the user")
			return
		}
		for _, v := range updateExamples {
			// Create the struct to update the user
			updateErr := auth.UpdateUserProfile(temp.UserId, v, *database)
			if updateErr == nil {
				t.Errorf("Error catching error in user profile ")
				return
			}
		}
	})
}
