package event

import (
	"fmt"
	"log"
	"time"

	"backend/pkg/auth"
	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func AllEventByGroup(groupId, userId string, database *structs.DB) ([]structs.Event, error) {
	var events []structs.Event
	var err error
	rows, err := database.DB.Query("SELECT * FROM Event WHERE groupId = '" + groupId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	for rows.Next() {
		var event structs.Event

		rows.Scan(&event.EventId, &event.UserId, &event.GroupId, &event.Name, &event.ImageUpload, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatedAt)

		EventP, _ := AllEventParticipant(event.EventId, database)

		for _, user := range EventP {
			if user.Status == 1 {
				event.NumOfParticipants++
				event.Participants = append([]structs.EventParticipant{user}, event.Participants...)
			}
		}

		events = append([]structs.Event{event}, events...)
	}
	return events, nil
}

func AddEventParticipant(eventId, userId string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	stmt, _ := database.DB.Prepare(`
	INSERT INTO EventParticipant values (?, ?, ?, ?)
`)

	_, err := stmt.Exec(eventId, userId, 0, createdAt)
	if err != nil {
		fmt.Println("inside Create Add Event Participant", err)
		return "", err
	}
	return eventId, nil
}

func GetEventParticipant(eventId, userId string, database *structs.DB) (structs.EventParticipant, error) {
	var eventParticipant structs.EventParticipant
	var err error
	rows, err := database.DB.Query("SELECT * FROM EventParticipant WHERE eventId = '" + eventId + "' AND userId = '" + userId + "'")
	if err != nil {
		fmt.Print(err)
		return eventParticipant, err
	}

	for rows.Next() {
		rows.Scan(&eventParticipant.EventId, &eventParticipant.UserId, &eventParticipant.Status, &eventParticipant.CreatedAt)
	}
	return eventParticipant, nil
}

func UpdateEventParticipant(event structs.Event, userId string, database structs.DB) (structs.Event, error) {
	participant, _ := GetEventParticipant(event.EventId, userId, &database)

	if participant.Status == 1 {
		update := "UPDATE EventParticipant SET status = ? WHERE eventId = '" + event.EventId + "' AND userId = '" + userId + "'"
		stmt, prepErr := database.DB.Prepare(update)
		if prepErr != nil {
			log.Println("Error updating field: ", prepErr)
			return event, prepErr
		}
		_, err := stmt.Exec(0)
		if err != nil {
			fmt.Println("Error updating ", err)
			return event, err
		}
		event.Status = "Not Going"

		var holder []structs.EventParticipant

		for _, user := range event.Participants {
			if user.UserId == userId {
				event.NumOfParticipants--
			} else {
				holder = append([]structs.EventParticipant{user}, holder...)
			}
		}
		event.Participants = holder
		return event, nil
	} else if participant.Status == 0 {

		update := "UPDATE EventParticipant SET status = ? WHERE eventId = '" + event.EventId + "' AND userId = '" + userId + "'"
		stmt, prepErr := database.DB.Prepare(update)
		if prepErr != nil {
			log.Println("Error updating field: ", prepErr)
			return event, prepErr
		}
		_, err := stmt.Exec(1)
		if err != nil {
			return event, err
		}
		event.Status = "Going"

		var holder structs.EventParticipant

		user, err := helper.GetUserInfo(userId, &database)
		if err != nil {
			return event, err
		}

		holder.Name = user.FirstName
		holder.UserId = userId
		holder.EventId = event.EventId
		holder.Status = 1

		event.NumOfParticipants++
		event.Participants = append([]structs.EventParticipant{holder}, event.Participants...)

		return event, nil
	}

	_, err := AddEventParticipant(event.EventId, userId, &database)
	if err != nil {
		return event, err
	}
	event.Status = "Not Going"
	return event, nil
}

func AllEventParticipant(eventId string, database *structs.DB) ([]structs.EventParticipant, error) {
	var eventParticipant structs.EventParticipant
	var eventParticipants []structs.EventParticipant
	var err error
	var reUser structs.User
	rows, err := database.DB.Query("SELECT * FROM EventParticipant WHERE eventId = '" + eventId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var status int
	var eventId2, userId, createdAt string
	for rows.Next() {
		rows.Scan(&eventId2, &userId, &status, &createdAt)

		err2 := auth.GetUser("userId", userId, &reUser, *database)
		if err2 != nil {
			fmt.Print("AllEventParticipant -", err2)
			return eventParticipants, err2
		}
		eventParticipant = structs.EventParticipant{
			EventId:   eventId2,
			UserId:    userId,
			Status:    status,
			Name:      reUser.FirstName,
			CreatedAt: createdAt,
		}

		eventParticipants = append([]structs.EventParticipant{eventParticipant}, eventParticipants...)

	}
	return eventParticipants, nil
}

func GetEventByEventId(eventId string, database *structs.DB) (structs.Event, error) {
	var event structs.Event
	var err error

	rows, err := database.DB.Query("SELECT * FROM Event WHERE eventId = '" + eventId + "'")
	if err != nil {
		fmt.Print(err)
		return event, err
	}
	var eventId2, userId, groupId, name, description, location, startTime, endTime, imageUpload string
	var createdAt time.Time
	for rows.Next() {
		rows.Scan(&eventId2, &userId, &groupId, &name, &imageUpload, &description, &location, &startTime, &endTime, &createdAt)
		event = structs.Event{
			EventId:     eventId2,
			UserId:      userId,
			GroupId:     groupId,
			Name:        name,
			Description: description,
			Location:    location,
			StartTime:   startTime,
			EndTime:     endTime,
			CreatedAt:   createdAt,
		}

	}
	return event, nil
}

func AllUserEvent(userId string, database *structs.DB) ([]structs.Event, error) {
	var eventStrA []string
	var events []structs.Event

	var err error
	rows, err := database.DB.Query("SELECT eventId FROM Event WHERE userId = '" + userId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var eventStr string
	for rows.Next() {
		rows.Scan(&eventStr)
		eventStrA = append([]string{eventStr}, eventStrA...)
	}

	for _, id := range eventStrA {
		userEvent, err := GetEventByEventId(id, database)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}
		events = append([]structs.Event{userEvent}, events...)
	}

	return events, nil
}

func AddEvent(groupId string, event structs.Event, database *structs.DB) (string, error) {
	createdAt := time.Now()
	eventId := uuid.NewV4().String()

	stmt, err := database.DB.Prepare(`
	INSERT INTO Event values (?, ?, ?, ?, ?, ?, ?, ?, ?,?)
`)
	if err != nil {
		l.LogMessage("Event.go", "AddEvent", err)
		return "", err
	}
	_, err = stmt.Exec(eventId, event.UserId, groupId, event.Name, "", event.Description, event.Location, event.StartTime, event.EndTime, createdAt)
	if err != nil {
		fmt.Println("inside Create Addevent", err)
		return "", err
	}
	err = CreateNewEventNotif(groupId, eventId, event.UserId, database)
	if err != nil {
		l.LogMessage("Event.go", "AddEvent", err)
		return "", err
	}

	return eventId, nil
}

// NotifMemberOfGroupEvent is a function that sends a notification to all members of a group when a new event is created
//
// Parameters:
//
//	groupId: the id of the group
//	eventId: the id of the event
//	database: the database
func CreateNewEventNotif(groupId, eventId, userId string, database *structs.DB) error {
	members, err := member.GetMembers(groupId, database)
	if err != nil {
		l.LogMessage("Event.go", "CreateNewEventNotif", err)
		return err
	}
	l.LogMessage("Event.go", "CreateNewEventNotif", len(members))
	for _, member := range members {
		l.LogMessage("Event.go", "CreateNewEventNotif", member.UserId)
		if member.UserId != userId {
			err := InsertEventNotification(eventId, member.UserId, database)
			if err != nil {
				l.LogMessage("Event.go", "CreateNewEventNotif", err)
				return err
			}
		}
	}
	return nil
}

// InsertEventNotification is a function that inserts a notification into the database
//
// Parameters:
//
//	eventId: the id of the event
//	userId: the id of the user
//	database: the database
func InsertEventNotification(eventId, userId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("INSERT INTO EventNotif values (?, ?, ?)")
	if err != nil {
		l.LogMessage("Event.go", "InsertEventNotification", err)
		return err
	}
	_, err = stmt.Exec(eventId, userId, 0)
	if err != nil {
		l.LogMessage("Event.go", "InsertEventNotification", err)
		return err
	}
	return nil
}

// GetEventNotifications is a function that gets all the notifications for a user
//
// Parameters:
//
//	userId: the id of the user
//	database: the database
func GetEventNotifications(userId string, database *structs.DB) ([]structs.GroupNotifWriter, error) {
	rows, err := database.DB.Query("SELECT * FROM EventNotif WHERE userId = '" + userId + "'")
	if err != nil {
		l.LogMessage("Event.go", "GetEventNotifications", err)
		return nil, err
	}
	var eventNotifcation structs.EventNotification
	var groupNotif structs.GroupNotifWriter
	var groupNotifs []structs.GroupNotifWriter
	for rows.Next() {
		rows.Scan(&eventNotifcation.EventId, &eventNotifcation.UserId, &eventNotifcation.Read)
		event, err := GetEventByEventId(eventNotifcation.EventId, database)
		if err != nil {
			return nil, err
		}
		group, err := helper.GetGroupInfo(event.GroupId, database)
		if err != nil {
			return nil, err
		}
		user, err := helper.GetUserInfo(event.UserId, database)
		if err != nil {
			return nil, err
		}
		groupNotif = structs.GroupNotifWriter{
			GroupId:   group,
			UserId:    user,
			EventId:   event,
			CreatedAt: event.CreatedAt,
			NotifType: "event",
			Read:      eventNotifcation.Read,
		}
		groupNotifs = append([]structs.GroupNotifWriter{groupNotif}, groupNotifs...)
	}
	return groupNotifs, nil
}
