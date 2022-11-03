package Test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/event"
	"backend/pkg/group"
	"backend/pkg/handler"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"
)

func TestHealthCheckEventHttpGet(t *testing.T) {
	database := DatabaseSetup()
	req := httptest.NewRequest(http.MethodGet, "/event", nil)
	w := httptest.NewRecorder()

	Env := handler.Env{Env: database}
	Env.Event(w, req)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestHealthCheckEventHttpPost(t *testing.T) {
	database := DatabaseSetup()
	req := httptest.NewRequest(http.MethodPost, "/event", nil)
	w := httptest.NewRecorder()

	Env := handler.Env{Env: database}
	Env.Event(w, req)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestCreateEvent(t *testing.T) {
	database := DatabaseSetup()
	t.Run("Creating a Event", func(t *testing.T) {
		group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: "wasfdfgfd"}
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
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: "wasfdfgfd2"}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group")
		}

		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: "UserId"}
		_, err := event.AddEvent(groupIdStr, event1, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		events, err := event.AllEventByGroup(groupIdStr, database)
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
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: "wasfdfgfd2"}
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
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: "wasfdfgfd2"}
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
		result, err := event.CheckIfUserInEventAndIfNotAddThem(addStr, "Ken", database)
		if err != nil {
			fmt.Println(err)
		}
		result2, err := event.CheckIfUserInEventAndIfNotAddThem(eventStr, "Ken2", database)
		if err != nil {
			fmt.Println(err)
		}
		want := true
		got := (result == false) && (result2 == true)
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
	t.Run("Create more event in the group and notif each user", func(t *testing.T) {
		// create a group
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: "wasfdfgfd2"}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group %v", errg)
		}
		// add member to the group
		_, err := member.AddMember(groupIdStr, "Ken", database)
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
		group1 := structs.Group{Name: "Pie2", Description: "Eating Pie2", Admin: "wasfdfgfd2"}
		groupIdStr, errg := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if errg != nil {
			t.Errorf("Error creating group %v", errg)
		}
		// add member to the group
		_, err = member.AddMember(groupIdStr, "Ken", database)
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
		// get all event notif
		notif, err := event.GetEventNotifications("Ken", database)
		l.LogMessage("Test", "Get all event notif", notif)
		l.LogMessage("Test", "Get all event notif", len(notif))
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		want := 1
		got := len(notif)
		if got != want {
			t.Errorf("Expected %v got %v", want, got)
		}
	})
}
