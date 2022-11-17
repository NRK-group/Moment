package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/closefriend"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestCloseFriends(t *testing.T) {
	Env := handler.Env{Env: database}

	closeEmailTwo := "follow@" + uuid.NewV4().String() + ".com"
	closeEmailOne := "follow@" + uuid.NewV4().String() + ".com"
	userOne := structs.User{FirstName: "Follow", LastName: "Follow", NickName: "Follow", AboutMe: "Follow", Email: closeEmailOne, Password: "Password123", DateOfBirth: "06-08-2002"}
	userTwo := structs.User{FirstName: "FollowTwo", LastName: "FollowTwo", NickName: "FollowTwo", AboutMe: "FollowTwo", Email: closeEmailTwo, Password: "Password123", DateOfBirth: "06-08-2002", IsPublic: 1}
	auth.InsertUser(userOne, *database)
	auth.InsertUser(userTwo, *database)
	var testOne, testTwo structs.User
	auth.GetUser("email", closeEmailOne, &testOne, *database)
	auth.GetUser("email", closeEmailTwo, &testTwo, *database)
	loginStruct := structs.User{Email: closeEmailOne, Password: "Password123"}
	loginUserBytes, err := json.Marshal(loginStruct)
	if err != nil {
		t.Errorf("Error marshalling the sampleUser")
	}

	// Create the bytes into a reader
	loginReq := bytes.NewReader(loginUserBytes)
	req := httptest.NewRequest(http.MethodPost, "/login", loginReq)

	w := httptest.NewRecorder()
	Env.Login(w, req)
	follow := structs.CloseFriend{
		UserId:        testOne.UserId,
		CloseFriendId: testTwo.UserId,
	}
	closeFriendBytes, err := json.Marshal(follow)
	if err != nil {
		t.Errorf("Error marshalling the sampleUser")
	}

	// Create the bytes into a reader
	closeFriend := bytes.NewReader(closeFriendBytes)
	closeFriendReq := httptest.NewRequest(http.MethodPost, "/closefriend", closeFriend)
	closeFriendReq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	wr := httptest.NewRecorder()
	Env.CloseFriends(wr, closeFriendReq)
	got := wr.Body.String()
	want := `{"Message":"Added"}`
	if got != want {
		t.Errorf("got %v. Want %v", got, want)
	}

	// Login again

	loginStruct2 := structs.User{Email: closeEmailOne, Password: "Password123"}
	loginUserBytes2, err := json.Marshal(loginStruct2)
	if err != nil {
		t.Errorf("Error marshalling the sampleUser")
	}
	// Create the bytes into a reader
	loginReq2 := bytes.NewReader(loginUserBytes2)
	req2 := httptest.NewRequest(http.MethodPost, "/login", loginReq2)

	w2 := httptest.NewRecorder()
	Env.Login(w2, req2)
	follow2 := structs.CloseFriend{
		UserId:        testOne.UserId,
		CloseFriendId: testTwo.UserId,
	}
	closeFriendBytes2, err := json.Marshal(follow2)
	if err != nil {
		t.Errorf("Error marshalling the sampleUser")
	}

	// Create the bytes into a reader
	closeFriend2 := bytes.NewReader(closeFriendBytes2)
	closeFriendReq2 := httptest.NewRequest(http.MethodPost, "/closefriend", closeFriend2)
	closeFriendReq2.Header = http.Header{"Cookie": w2.Header()["Set-Cookie"]}
	wr2 := httptest.NewRecorder()
	Env.CloseFriends(wr2, closeFriendReq2)
	got2 := wr2.Body.String()
	want2 := `{"Message":"Removed"}`
	if got2 != want2 {
		t.Errorf("got %v. Want %v", got2, want2)
	}
}

func TestCloseFriendList(t *testing.T) {
	Env := handler.Env{Env: database}
	profileUser := CreateUser(database, t)
	var ids []string
	for i:=0; i<3; i++ {
		//Create a user and make them follow profile user
		tempUser := CreateUser(database, t)
		closefriend.AddCloseFriend(profileUser.UserId, tempUser.UserId, *database)
		ids = append(ids, tempUser.UserId)
	}
	loginStruct := structs.User{Email: profileUser.Email, Password: "Password123"}
	logUser, _ := json.Marshal(loginStruct)
	// Create the bytes into a reader
	logReq := bytes.NewReader(logUser)
	req := httptest.NewRequest(http.MethodPost, "/login", logReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	getReq := httptest.NewRequest(http.MethodGet, "/getclosefriend", nil)
	getReq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	wr := httptest.NewRecorder()
	Env.CloseFriendList(wr, getReq)
	var got []structs.Info
	json.Unmarshal(wr.Body.Bytes(), &got)
	if len(got) != 3 {
		t.Errorf("Expected length 3 got %v", len(got))
	}
	for i, v := range got {
		if v.Id != ids[i] {
			t.Errorf("got %v. Want %v.", v.Id, ids[i])
		}
	}
}