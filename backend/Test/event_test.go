package Test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/event"
	"backend/pkg/group"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"
)

func TestHealthCheckEventHttpGet(t *testing.T) {
	w, Env, _ := LoginUser(database, t)
	reqq := httptest.NewRequest(http.MethodGet, "/event", nil)
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	Env.Event(w, reqq)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestHealthCheckEventHttpPost(t *testing.T) {
	w, Env, _ := LoginUser(database, t)
	reqq := httptest.NewRequest(http.MethodPost, "/event", nil)
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}

	Env.Event(w, reqq)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestCreateEvent(t *testing.T) {
	database := DatabaseSetup()
	t.Run("Creating a Event", func(t *testing.T) {
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: newUser.UserId}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group")
		}
		event1 := structs.Event{Name: "Test1", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId3"}
		str, err := event.AddEvent(groupIdStr, event1, database)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	t.Run("Get all Events of a group", func(t *testing.T) {
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: newUser.UserId}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group")
		}

		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId"}
		_, err := event.AddEvent(groupIdStr, event1, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		events, err := event.AllEventByGroup(groupIdStr, newUser.UserId, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		want := groupIdStr
		got := events[0].GroupId
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})

	t.Run("Add user to Events", func(t *testing.T) {
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: newUser.UserId}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group")
		}
		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId"}
		eventStr, err := event.AddEvent(groupIdStr, event1, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		addStr, err := event.AddEventParticipant(eventStr, "Ken", database)
		if err != nil {
			t.Errorf("Add Event Participant")
		}
		fmt.Println(addStr, " -- ", eventStr)
		want := addStr
		got := eventStr
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})

	t.Run("Add user that already in a Events", func(t *testing.T) {
		newUser := CreateUser(database, t)
		newUser2 := CreateUser(database, t)
		newUser3 := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: newUser.UserId}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group")
		}
		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId"}
		eventStr, err := event.AddEvent(groupIdStr, event1, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		eventId, err := event.AddEventParticipant(eventStr, newUser2.UserId, database)
		if err != nil {
			t.Errorf("Add Event Participant")
		}
		result, err := event.GetEventParticipant(eventStr, newUser2.UserId, database)
		if err != nil {
			fmt.Println(err)
		}
		result2, err := event.GetEventParticipant(eventStr, newUser3.UserId, database)
		if err != nil {
			fmt.Println(err)
		}
		want := true
		got := (result.EventId == eventId) && (result2.EventId == "")
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Create more event in the group and notif each user", func(t *testing.T) {
		// create a group
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: newUser.UserId}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group %v", errg)
		}
		// add member to the group
		_, err := member.AddMember(groupIdStr, newUser.UserId, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		// add event to the group
		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId"}
		_, err = event.AddEvent(groupIdStr, event1, database)
		l.LogMessage("Test", "Create more event in the group and notif each user", err)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	t.Run("Get all event notif", func(t *testing.T) {
		// create a group
		// Delete all the ad userid on eventnotif
		_, err := database.DB.Exec("DELETE FROM eventnotif WHERE userid = 'Ken'")
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		newUser3 := CreateUser(database, t)
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: newUser3.UserId}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group %v", errg)
		}
		// add member to the group
		_, err = member.AddMember(groupIdStr, newUser.UserId, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		// add event to the group
		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId"}
		_, err = event.AddEvent(groupIdStr, event1, database)
		l.LogMessage("Test", "Create more event in the group and notif each user", err)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		user := CreateUser(database, t)
		groupId, err := group.CreateGroup("Pie3", "Eating Pie3", user.UserId, database)
		if err != nil {
			t.Errorf("Error creating group %v", err)
		}
		user2 := CreateUser(database, t)
		member.AddMember(groupId, user2.UserId, database)
		event2 := structs.Event{UserId: user.UserId, GroupId: groupId, Name: "Test5", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime"}
		eventId, err := event.AddEvent(groupId, event2, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		notif, err := event.GetEventNotifications(user2.UserId, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		if eventId != notif[0].EventId.EventId {
			t.Errorf("Expected %s got %s", eventId, notif[0].EventId.EventId)
		}
		want := 1
		got := len(notif)
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}
