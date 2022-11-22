package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/event"
	"backend/pkg/follow"
	"backend/pkg/group"
	"backend/pkg/handler"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"
)

func TestNotification(t *testing.T) {
	Env := handler.Env{Env: database}
	user1 := CreateUser(database, t)
	user2 := CreateUser(database, t)
	follow.FollowUser(user2.UserId, user1.UserId, database)
	group1, err := group.CreateGroup("TestGroup", "TestGroup", user1.UserId, database)
	if err != nil {
		t.Errorf("Error creating group1: %v", err)
	}
	member.AddInvitationNotif(group1, user2.UserId, user1.UserId, "join", database)
	group2, err := group.CreateGroup("Second Group", "Coding", user2.UserId, database)
	if err != nil {
		t.Errorf("Error creating group2: %v", err)
	}
	member.AddMember(group2, user1.UserId, database)
	event1 := structs.Event{UserId: user2.UserId, GroupId: group2, Name: "Nate's Event", Description: "Coding", Location: "London", StartTime: "StartTime", EndTime: " EndTime"}
	_, err = event.AddEvent(group2, event1, database)
	if err != nil {
		t.Errorf("Error creating event %v", err)
	}
	// event.AddEvent()
	sampleUser := &structs.User{
		Email: user1.Email, Password: "Password123",
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
		if followNotif[0].UserId.Id != user2.UserId {
			t.Errorf("Expected body %s, got %q", user2.UserId, followNotif[0].FollowingId.Id)
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
		if len(groupNotif) != 2 {
			t.Errorf("Expected body %s, got %d", "2", len(groupNotif))
		}
		if groupNotif[0].GroupId.Id != group1 {
			t.Error("Expected ", group1, " got ", groupNotif[0].GroupId.Id)
		}
		if groupNotif[0].ReceiverId.Id != user1.UserId {
			t.Error("Expected ", user1.UserId, " got ", groupNotif[0].ReceiverId.Id)
		}
		if groupNotif[0].UserId.Id != user2.UserId {
			t.Error("Expected ", user2.UserId, " got ", groupNotif[0].UserId.Id)
		}
		if groupNotif[1].NotifType != "event" {
			t.Errorf("Expected body %s, got %s", "event", groupNotif[1].NotifType)
		}
		if groupNotif[1].GroupId.Id != group2 {
			t.Errorf("Expected body %s, got %s", group2, groupNotif[1].GroupId.Id)
		}
	})
}
