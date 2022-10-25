package Test

import (
	"fmt"
	"strconv"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

var tests = []structs.User {
	{FirstName: "", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "", Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "0000-00-00", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},

	{FirstName: "Length", LastName: "Length", NickName: "Length", Email: "Length"+uuid.NewV4().String(), Password: "Length",
	DateOfBirth: "0000-00-00", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
	IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,},
}

func TestInsertUser(t *testing.T) {
	randEmail := "insertUSer@"+uuid.NewV4().String()
	t.Run("Insert valid user to DB", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		sampleUser := &structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *DB)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		//Query the db to check if the user was inserted
		rows, err := DB.DB.Query(`SELECT * FROM User WHERE Email = ?`, sampleUser.Email)
		var userId, sessionId, firstName, lastName, nickName, email, DOB, avatar, aboutMe, createdAt, isLoggedIn, isPublic, numFollowers, numFollowing, numPosts, password string
		var resultUser *structs.User

		for rows.Next() {
			rows.Scan(&userId, &sessionId, &firstName, &lastName, &nickName, &email, &DOB, &avatar, &aboutMe, &createdAt, &isLoggedIn, &isPublic, &numFollowers, &numFollowing, &numPosts, &password)

			resultUser = &structs.User{
				UserId:      "-",
				SessionId:   sessionId,
				FirstName:   firstName,
				LastName:    lastName,
				NickName:    nickName,
				Email:       email,
				DateOfBirth: DOB,
				Avatar:      avatar,
				AboutMe:     aboutMe,
				CreatedAt:   "-",
				Password:    password,
			}
		}

		sampleUser.Password = strconv.FormatBool(auth.CheckPasswordHash(sampleUser.Password, resultUser.Password))
		if err != nil {
			t.Errorf("Error hashing the password %v", err)
		}
		resultUser.Password = "true"
		want := sampleUser
		got := resultUser

		if *got != *want {
			t.Errorf("want %v, \n got %v", want, got)
		}
	})
	t.Run("inserting a user with used email to the db", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		sampleUser := &structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *DB)
		if err == nil {
			t.Errorf("Error Catching already used email %v", err)
		}
	})
	t.Run("Check the length of neccesary values can't be 0", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")
		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}

		for i, s := range tests {
			err := auth.InsertUser(s, *DB)
			fmt.Println("Index:", i, "CURRENT ", s)
			fmt.Println()
		if err == nil {
			t.Errorf("Error Catching empty values %v", err)
		}
		}
	})
}

func TestCheckCredentials(t *testing.T) {
	testEmail := "GetUser@" + uuid.NewV4().String()

	t.Run("Non-existing account entered", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}

		sampleUser := &structs.User{
			FirstName: "GetUser", LastName: "GetUser", NickName: "GetUser", Email: "123", Password: "GetUser",
			DateOfBirth: "GetUser", AboutMe: "GetUser", Avatar: "GetUser", CreatedAt: "GetUser", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		//Now check if the value is input to the db
		wantStr := "Account not found"
		wantBool := false

		gotBool, gotStr := auth.CheckCredentials(sampleUser.Email, sampleUser.Password, DB)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}
	})

	t.Run("Check with correct credentials", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}
		// Env := &handler.Env{Env: DB}

		sampleUser := &structs.User{
			FirstName: "GetUser", LastName: "GetUser", NickName: "GetUser", Email: testEmail, Password: "GetUser123",
			DateOfBirth: "GetUser", AboutMe: "GetUser", Avatar: "GetUser", CreatedAt: "GetUser", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}

		err := auth.InsertUser(*sampleUser, *DB)

		if err != nil {
			t.Errorf("Error Inserting the test user to the db")
		}

		//Now check if the value is input to the db
		wantStr := "Valid Login"
		wantBool := true

		gotBool, gotStr := auth.CheckCredentials(sampleUser.Email, sampleUser.Password, DB)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}
	})

	t.Run("Check with incorrect password", func(t *testing.T) {
		// Create the database that will be used for testing
		database := sqlite.CreateDatabase("./social_network_test.db")

		// migrate the database
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")

		// Create the database struct
		DB := &structs.DB{DB: database}

		//Now check if the value is input to the db
		wantStr := "Incorrect Password"
		wantBool := false

		gotBool, gotStr := auth.CheckCredentials(testEmail, "incorrectPassword", DB)

		if gotBool != wantBool && gotStr != wantStr {
			t.Errorf("Got: %v %v, Want: %v %v", gotBool, gotStr, wantBool, wantStr)
		}
	})
}

func TestGetUser(t *testing.T){
	randEmail:= uuid.NewV4().String()
	t.Run("Getting valid user", func(t *testing.T) {
		database := sqlite.CreateDatabase("./social_network_test.db")
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")// migrate the database
		DB := &structs.DB{DB: database}// Create the database struct

		sampleUser := &structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: randEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *DB)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		var got structs.User
		getErr := auth.GetUser(`email`, randEmail, &got, *DB)
		if getErr != nil {
			t.Errorf("Error getting the user from the database")
		}
		got.CreatedAt="-"
		got.UserId="-"
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
		database := sqlite.CreateDatabase("./social_network_test.db")
		sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")// migrate the database
		DB := &structs.DB{DB: database}// Create the database struct
		testEmail := "random@" + uuid.NewV4().String()
		sampleUser := &structs.User{
			FirstName: "InsertUser", LastName: "InsertUser", NickName: "InsertUser", Email: testEmail, Password: "InsertUser",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "-", UserId: "-", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		err := auth.InsertUser(*sampleUser, *DB)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		var got structs.User
		getErr := auth.GetUser("email", testEmail, &got, *DB)
		if getErr != nil {
			t.Errorf("Error getting the user from the database")
		}
		got.CreatedAt="-"
		got.UserId="-"
		sampleUser.Password = strconv.FormatBool(auth.CheckPasswordHash(sampleUser.Password, got.Password))
		if err != nil {
			t.Errorf("Error hashing the password %v", err)
		}
		got.Password = "true"

		if got != *sampleUser {
			t.Errorf("got: %v. Want: %v.", got, *sampleUser)
		}
	})
}