package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// TestFollow will include all the tests for the follow feature
func TestFollow(t *testing.T) {
	// Create the database that will be used for testing
	followerId := uuid.NewV4().String()
	followingId := uuid.NewV4().String()
	t.Run("Insert follow", func(t *testing.T) {
		got := follow.InsertFollow(followerId, followingId, database)
		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})
	followerId = "8a5e7163-7cb7-440a-9537-42a9c4371752"
	followingId = "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2"
	t.Run("Delete follow", func(t *testing.T) {
		err := follow.DeleteFollow(followerId, followingId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
	})
	t.Run("Check follow false", func(t *testing.T) {
		got := follow.CheckIfFollow(followerId, followingId, database)
		want := false
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
	t.Run("Check follow true", func(t *testing.T) {
		follow.InsertFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		got := follow.CheckIfFollow("8a5e7163-7cb7-440a-9537-42a9c4371752", "9df2d6fb-0ecc-43da-a82e-fd4f4fac51f2", database)
		want := true
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
	// Create the users that will be used for testing
	currentEmail := "hello" + uuid.NewV4().String() + "@test.com"
	currentUser := &structs.User{
		FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: currentEmail, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*currentUser, *database)
	var currentResult structs.User
	auth.GetUser("email", currentEmail, &currentResult, *database)
	t.Run("Follow user returns follow", func(t *testing.T) {
		// Create the users that will be used for testing
		email := "hello" + uuid.NewV4().String() + "@test.com"
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		status, err := follow.FollowUser(currentResult.UserId, result.UserId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := status
		want := "follow"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Follow user returns pending", func(t *testing.T) {
		// Create the users that will be used for testing
		email := "hello" + uuid.NewV4().String() + "@test.com"
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		status, err := follow.FollowUser(currentResult.UserId, result.UserId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := status
		want := "pending"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Follow user returns unfollow", func(t *testing.T) {
		// Create the users that will be used for testing
		email := "hello" + uuid.NewV4().String() + "@test.com"
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		follow.InsertFollow(currentResult.UserId, result.UserId, database)
		status, err := follow.FollowUser(currentResult.UserId, result.UserId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := status
		want := "unfollow"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Insert follow notif", func(t *testing.T) {
		// Create the users that will be used for testing
		email := "hello" + uuid.NewV4().String() + "@test.com"
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		err := follow.InsertFollowNotif(currentResult.UserId, result.UserId, "pending", database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
	})
	t.Run("Get num of following", func(t *testing.T) {
		// Create the users that will be used for testing
		email := "hello" + uuid.NewV4().String() + "@test.com"
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
		}
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		num, err := follow.GetNumOfFollowing(result.UserId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := num
		want := 5
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Get num of following", func(t *testing.T) {
		// Create the users that will be used for testing
		email := "hello" + uuid.NewV4().String() + "@test.com"
		inputUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
		}
		auth.InsertUser(*inputUser, *database)
		var result structs.User
		auth.GetUser("email", email, &result, *database)
		num, err := follow.GetNumOfFollowers(result.UserId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := num
		want := 5
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Get following notifs", func(t *testing.T) {
		// Create the users that will be used for testing
		user := CreateUser(database, t)
		user2 := CreateUser(database, t)
		follow.InsertFollowNotif(user.UserId, user2.UserId, "pending", database)
		fot, err := follow.GetFollowNotifs(user2.UserId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := len(fot)
		want := 1
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Pending deleted", func(t *testing.T) {
		emailOne := "follow@" + uuid.NewV4().String() + ".com"
		emailTwo := "follow@" + uuid.NewV4().String() + ".com"
		Env := handler.Env{Env: database}

		userOne := structs.User{FirstName: "Follow", LastName: "Follow", NickName: "Follow", AboutMe: "Follow", Email: emailOne, Password: "Password123", DateOfBirth: "06-08-2002"}
		userTwo := structs.User{FirstName: "FollowTwo", LastName: "FollowTwo", NickName: "FollowTwo", AboutMe: "FollowTwo", Email: emailTwo, Password: "Password123", DateOfBirth: "06-08-2002", IsPublic: 0}
		auth.InsertUser(userOne, *database)
		auth.InsertUser(userTwo, *database)
		var testOne, testTwo structs.User
		auth.GetUser("email", emailOne, &testOne, *database)
		auth.GetUser("email", emailTwo, &testTwo, *database)
		// login userone
		loginOne := structs.User{Email: emailOne, Password: "Password123"}
		loginBytes, err := json.Marshal(loginOne)
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
			return
		}

		// Create the bytes into a reader
		testlogin := bytes.NewReader(loginBytes)
		req := httptest.NewRequest(http.MethodPost, "/login", testlogin)
		w := httptest.NewRecorder()
		Env.Login(w, req)
		// Now let user one follow user two
		followPending := structs.Follower{FollowerId: testOne.UserId, FollowingId: testTwo.UserId}
		followReqBytes, _ := json.Marshal(followPending)
		followReader := bytes.NewReader(followReqBytes)
		follow.FollowUser(testOne.UserId, testTwo.UserId, database)

		reqFol := httptest.NewRequest(http.MethodPut, "/followrequest", followReader)
		reqFol.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
		wr := httptest.NewRecorder()
		Env.FollowReq(wr, reqFol)
		got := wr.Body.String()
		want := `{"Message":"unfollow"}`
		if got != want {
			t.Errorf("got %v \n want %v", got, want)
		}
	})
}

func TestDeletePendingRequests(t *testing.T) {
	// Create three users
	one := CreateUser(database, t)
	two := CreateUser(database, t)
	three := CreateUser(database, t)
	follow.InsertFollowNotif(two.UserId, one.UserId, "pending", database)
	follow.InsertFollowNotif(three.UserId, one.UserId, "pending", database)

	follow.DeletePendingRequests(one.UserId, *database)

	rows, _ := database.DB.Query("SELECT * FROM FollowNotif WHERE followingId = ?", one.UserId)
	defer rows.Close()
	got := 0
	for rows.Next() {
		got++
	}
	want := 0

	if got != want {
		t.Errorf("got %v \n want %v", got, want)
	}

	err := follow.DeletePendingRequests(one.UserId, *database)
	if err != nil {
		t.Errorf("got %v \n want %v", err, nil)
	}

}
