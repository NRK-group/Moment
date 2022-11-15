package Test

import (
	"testing"
	"time"

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
		time.Sleep(time.Second * 5)
		want := v
		if got != want {
			t.Errorf("Got %v . Want %v .", got, want)
		}
	}
}
