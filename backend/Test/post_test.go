package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/db/sqlite"
	"backend/pkg/handler"
	"backend/pkg/post"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func DatabaseSetup() *structs.DB {
	// this open or create the database
	db := sqlite.CreateDatabase("./social_network_test.db")
	// migrate the database
	sqlite.MigrateDatabase("file://../pkg/db/migrations/sqlite", "sqlite3://./social_network_test.db")
	// initialize the database struct
	database := &structs.DB{DB: db}

	return database
}

func CreateUser(database *structs.DB, t *testing.T) structs.User {
	currentEmail := "hello" + uuid.NewV4().String() + "@test.com"
	currentUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: currentEmail, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*currentUser, *database)
	var currentResult structs.User
	auth.GetUser("email", currentEmail, &currentResult, *database)
	return currentResult
}

func CreatePost(GroupId string, database *structs.DB, t *testing.T) string {
	newUser := CreateUser(database, t)
	post1 := structs.Post{UserID: newUser.UserId, Content: "hey", GroupID: GroupId, Image: "wasfdfgfd"}
	postId, err := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, post1.Privacy, database)
	if err != nil {
		t.Errorf("Error Inserting the struct into the db %v", err)
	}
	return postId
}

func createNewPost(userId string, database *structs.DB) {
	post1 := structs.Post{UserID: userId, Content: "hey" + uuid.NewV4().String(), Privacy: 0}
	_, err := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, post1.Privacy, database)
	if err != nil {
		fmt.Println("Error Inserting the struct into the db", err.Error())
	}
}

func TestHealthCheckPostHandlerHttpGet(t *testing.T) {
	w, Env, _ := LoginUser(database, t)
	// database := DatabaseSetup()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqq, err := http.NewRequest("GET", "/post", nil)
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Post)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, reqq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected2 := http.StatusOK
	if status := rr.Code; status != expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected2)
	}
}

func TestCreatePost(t *testing.T) {
	t.Run("Insert Post to DB", func(t *testing.T) {
		// database := DatabaseSetup()
		newUser := CreateUser(database, t)
		post1 := structs.Post{UserID: newUser.UserId, Content: "hey", GroupID: "", Image: "wasfdfgfd"}
		_, err := post.CreatePost(post1.UserID, post1.GroupID, post1.Image, post1.Content, 0, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Read all Posts from the DB", func(t *testing.T) {
		// database := DatabaseSetup()
		posts, err := post.AllPost(database)
		if err != nil || len(posts) == 0 {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}

func TestPostHandlerMakeAPost(t *testing.T) {
	// database := DatabaseSetup()
	w, Env, _ := LoginUser(database, t)
	reqq := httptest.NewRequest(http.MethodGet, "/event", nil)
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}

	newUser := CreateUser(database, t)
	post1 := structs.Post{UserID: newUser.UserId, Content: "hey", Image: "wasfdfgfd"}
	body, _ := json.Marshal(post1)

	reqq, err := http.NewRequest("POST", "/post", bytes.NewBuffer(body))
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Post)
	handler.ServeHTTP(rr, reqq)
	expected := rr.Code
	expectedStr := 200
	if expectedStr != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			expected, expectedStr)
	}
}

func TestPostHandlerGettingAllPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/post", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Env := &handler.Env{Env: database}
	handler := http.HandlerFunc(Env.Post)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	expected := http.StatusOK
	if status := rr.Code; status != expected && strings.Contains(rr.Body.String(), "PostID") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

var newUser, newUser2 structs.User

func TestGetingUserPosts(t *testing.T) {
	t.Run("Insert Posts to DB", func(t *testing.T) {
		newUser = CreateUser(database, t)
		newUser2 = CreateUser(database, t)
		createNewPost(newUser2.UserId, database)
		createNewPost(newUser2.UserId, database)
		createNewPost(newUser2.UserId, database)
		createNewPost(newUser.UserId, database)
	})

	t.Run("Read all user Posts from the DB", func(t *testing.T) {
		posts, err := post.AllUserPost(newUser.UserId, database)
		if err != nil || len(posts) != 1 {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}

func TestGetUserPostHandlerGetAllUserPost(t *testing.T) {
	t.Run("Creating users and Posts to DB", func(t *testing.T) {
		newUser = CreateUser(database, t)
		newUser2 = CreateUser(database, t)
		createNewPost(newUser2.UserId, database)
		createNewPost(newUser2.UserId, database)
		createNewPost(newUser2.UserId, database)
		createNewPost(newUser.UserId, database)
	})

	t.Run("Testing the get user post handler ", func(t *testing.T) {
		loginStruct := structs.User{Email: newUser.Email, Password: "Password123"}

		loginUserBytes, err := json.Marshal(loginStruct)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
		Env := &handler.Env{Env: database}

		// Create the bytes into a reader
		loginReq := bytes.NewReader(loginUserBytes)
		req := httptest.NewRequest(http.MethodPost, "/login", loginReq)
		w := httptest.NewRecorder()
		Env.Login(w, req)

		reqGetUserPosts, err := http.NewRequest("GET", "/getUserPosts", nil)
		reqGetUserPosts.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
		values := reqGetUserPosts.URL.Query()
		values.Add("userID", newUser.UserId)
		reqGetUserPosts.URL.RawQuery = values.Encode()
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Env.GetUserPosts)
		fmt.Println("test")
		handler.ServeHTTP(rr, reqGetUserPosts)
		expected := http.StatusOK
		if status := rr.Code; status != expected && len(strings.Split(rr.Body.String(), "PostID")) == 2 {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
