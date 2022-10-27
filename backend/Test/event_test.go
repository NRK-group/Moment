package Test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/event"
	"backend/pkg/group"
	"backend/pkg/handler"
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
	t.Run("Creating a Event", func(t *testing.T) {
		database := DatabaseSetup()

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
		database := DatabaseSetup()
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
		database := DatabaseSetup()
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
}
