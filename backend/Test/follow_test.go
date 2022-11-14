package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/follow"
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
		FollowingTestId := uuid.NewV4().String()
		follow.InsertFollowNotif("Hello1", FollowingTestId, "pending", database)
		follow.InsertFollowNotif("Hello2", FollowingTestId, "pending", database)
		follow.InsertFollowNotif("Hello3", FollowingTestId, "pending", database)
		fot, err := follow.GetFollowingNotifs(FollowingTestId, database)
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
		got := len(fot)
		want := 3
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

// func TestCheckIfFollowPending(t *testing.T) {
// 	t.Run("Follow is pending", func(t *testing.T) {
// 		Env := handler.Env{Env: database}
// 		emailOne := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
// 		emailTwo := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")

// 		// Create the struct that will be inserted
// 		firstUser := structs.User{
// 			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailOne, Password: "TestPass",
// 			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
// 			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
// 		}
// 		auth.InsertUser(firstUser, *Env.Env)


// 		// Create the struct that will be inserted
// 		secondUser := structs.User{
// 			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailTwo, Password: "TestPass",
// 			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
// 			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
// 		}

// 		auth.InsertUser(secondUser, *Env.Env)
// 		// Get the two users
// 		var userOne, userTwo structs.User
// 		auth.GetUser("email", emailOne, &userOne, *Env.Env)
// 		auth.GetUser("email", emailTwo, &userTwo, *Env.Env)
// 		follow.InsertFollowNotif(userOne.UserId, userTwo.UserId, "pending", database)
// 		got := follow.CheckIfFollowPending(userOne.UserId, userTwo.UserId, database)
// 		want := true
// 		if got != want {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// 	t.Run("follow doesn't equal pending", func(t *testing.T) {
// 		Env := handler.Env{Env: database}
// 		emailOne := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
// 		emailTwo := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")

// 		// Create the struct that will be inserted
// 		firstUser := structs.User{
// 			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailOne, Password: "TestPass",
// 			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
// 			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
// 		}
// 		auth.InsertUser(firstUser, *Env.Env)


// 		// Create the struct that will be inserted
// 		secondUser := structs.User{
// 			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailTwo, Password: "TestPass",
// 			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
// 			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
// 		}

// 		auth.InsertUser(secondUser, *Env.Env)
// 		// Get the two users
// 		var userOne, userTwo structs.User
// 		auth.GetUser("email", emailOne, &userOne, *Env.Env)
// 		auth.GetUser("email", emailTwo, &userTwo, *Env.Env)
// 		follow.InsertFollowNotif(userOne.UserId, userTwo.UserId, "follow", database)
// 		got := follow.CheckIfFollowPending(userOne.UserId, userTwo.UserId, database)
// 		want := true
// 		if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// 	t.Run("Not row in the db", func(t *testing.T) {
// 		Env := handler.Env{Env: database}
// 		emailOne := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")
// 		emailTwo := strings.ToLower("handlertest@" + uuid.NewV4().String() + ".com")

// 		// Create the struct that will be inserted
// 		firstUser := structs.User{
// 			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailOne, Password: "TestPass",
// 			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
// 			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
// 		}
// 		auth.InsertUser(firstUser, *Env.Env)


// 		// Create the struct that will be inserted
// 		secondUser := structs.User{
// 			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: emailTwo, Password: "TestPass",
// 			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
// 			IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
// 		}

// 		auth.InsertUser(secondUser, *Env.Env)
// 		// Get the two users
// 		var userOne, userTwo structs.User
// 		auth.GetUser("email", emailOne, &userOne, *Env.Env)
// 		auth.GetUser("email", emailTwo, &userTwo, *Env.Env)
// 		got := follow.CheckIfFollowPending(userOne.UserId, userTwo.UserId, database)
// 		want := false
// 		if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }