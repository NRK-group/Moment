package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/group"
	"backend/pkg/handler"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestNotification(t *testing.T) {
	Env := handler.Env{Env: database}
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	email2 := "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email2, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	auth.InsertUser(*user2, *database)
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *database)
	follow.FollowUser(result2.UserId, result1.UserId, database)
	group, _ := group.CreateGroup("TestGroup", "TestGroup", result1.UserId, database)
	member.AddInvitationNotif(group, result2.UserId, result1.UserId, "join", database)
	sampleUser := &structs.User{
		Email: email, Password: "Password123",
	}
	sampleUserBytes, err := json.Marshal(sampleUser)
	t.Run("Test Marshal data", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
	})
	testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
	req := httptest.NewRequest(http.MethodPost, "/login", testReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	t.Run("Test login", func(t *testing.T) {
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	req2 := httptest.NewRequest(http.MethodGet, "/notification", nil)
	req2.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values := req2.URL.Query()
	values.Add("notifType", "follow")
	req2.URL.RawQuery = values.Encode()
	w2 := httptest.NewRecorder()
	Env.Notification(w2, req2)
	t.Run("Test follow notification", func(t *testing.T) {
		if w2.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w2.Code)
		}
		var followNotif []structs.FollowerNotif
		err := json.NewDecoder(w2.Body).Decode(&followNotif)
		if err != nil {
			t.Errorf("Error decoding the response body %s", err)
		}
		if len(followNotif) != 1 {
			t.Errorf("Expected body %s, got %d", "1", len(followNotif))
		}
		if followNotif[0].UserId.Id != result2.UserId {
			t.Errorf("Expected body %s, got %q", result2.UserId, followNotif[0].FollowingId.Id)
		}
	})
	req3 := httptest.NewRequest(http.MethodGet, "/notification", nil)
	req3.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values = req3.URL.Query()
	values.Add("notifType", "general")
	req3.URL.RawQuery = values.Encode()
	w3 := httptest.NewRecorder()
	Env.Notification(w3, req3)
	t.Run("Test general notification", func(t *testing.T) {
		if w3.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w3.Code)
		}
		var generalNotif string
		err = json.NewDecoder(w3.Body).Decode(&generalNotif)
		if generalNotif != "general" {
			t.Errorf("Expected body %s, got %s", "general", w3.Body)
		}
	})
	req4 := httptest.NewRequest(http.MethodGet, "/notification", nil)
	req4.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values = req4.URL.Query()
	values.Add("notifType", "group")
	req4.URL.RawQuery = values.Encode()
	w4 := httptest.NewRecorder()
	Env.Notification(w4, req4)
	t.Run("Test group notification", func(t *testing.T) {
		if w4.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w4.Code)
		}
		var groupNotif []structs.GroupNotifWriter
		err = json.NewDecoder(w4.Body).Decode(&groupNotif)
		l.LogMessage("Tes", "decode", groupNotif)
		if len(groupNotif) != 1 {
			t.Errorf("Expected body %s, got %d", "1", len(groupNotif))
		}
		if groupNotif[0].GroupId.Id != group {
			t.Error("Expected ", group, " got ", groupNotif[0].GroupId.Id)
		}
		if groupNotif[0].UserId.Id != result2.UserId {
			t.Error("Expected ", result1.UserId, " got ", groupNotif[0].UserId.Id)
		}
		if groupNotif[0].ReceiverId.Id != result1.UserId {
			t.Error("Expected ", result2.UserId, " got ", groupNotif[0].ReceiverId.Id)
		}
	})
}
