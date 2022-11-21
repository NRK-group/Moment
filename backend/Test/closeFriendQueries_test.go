package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/closefriend"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

var results = []string{"Added", "Removed"}

func TestUpdateCloseFriend(t *testing.T) {
	// Create two usersemailOne := "follow@" + uuid.NewV4().String() + ".com"
	closeEmailTwo := "follow@" + uuid.NewV4().String() + ".com"
	closeEmailOne := "follow@" + uuid.NewV4().String() + ".com"
	userOne := structs.User{FirstName: "Follow", LastName: "Follow", NickName: "Follow", AboutMe: "Follow", Email: closeEmailOne, Password: "Password123", DateOfBirth: "06-08-2002"}
	userTwo := structs.User{FirstName: "FollowTwo", LastName: "FollowTwo", NickName: "FollowTwo", AboutMe: "FollowTwo", Email: closeEmailTwo, Password: "Password123", DateOfBirth: "06-08-2002", IsPublic: 1}
	auth.InsertUser(userOne, *database)
	auth.InsertUser(userTwo, *database)
	var testOne, testTwo structs.User
	auth.GetUser("email", closeEmailOne, &testOne, *database)
	auth.GetUser("email", closeEmailTwo, &testTwo, *database)
	for _, v := range results {
		got := closefriend.UpdateCloseFriend(testOne.UserId, testTwo.UserId, *database)
		want := v
		if got != want {
			t.Errorf("Got %v . Want %v .", got, want)
		}
	}
}

func TestGetCloseFriends(t *testing.T) {
	profileUser := CreateUser(database, t)
	var ids []string
	for i := 0; i < 3; i++ {
		// Create a user and make them follow profile user
		tempUser := CreateUser(database, t)
		closefriend.AddCloseFriend(profileUser.UserId, tempUser.UserId, *database)
		ids = append(ids, tempUser.UserId)
	}
	got := closefriend.GetCloseFriends(profileUser.UserId, *database)
	if len(got) != 3 {
		t.Errorf("Expected length 3 got %v", len(got))
	}
	for i, v := range got {
		if v.Id != ids[i] {
			t.Errorf("got %v. Want %v.", v.Id, ids[i])
		}
	}
}

func TestCurrentCloseFriend(t *testing.T) {
	// Create two users
	userOne := CreateUser(database, t)
	userTwo := CreateUser(database, t)
	got := closefriend.CurrentCloseFriend(userOne.UserId, userTwo.UserId, *database)
	want := false
	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
		return
	}
	closefriend.AddCloseFriend(userOne.UserId, userTwo.UserId, *database)
	got = closefriend.CurrentCloseFriend(userOne.UserId, userTwo.UserId, *database)
	want = true
	if got != want {
		t.Errorf("got: %v, want: %v.", got, want)
		return
	}
}
