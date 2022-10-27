package event

import (
	"fmt"

	"backend/pkg/structs"
)

func AllEventByGroup(groupId string, database *structs.DB) ([]structs.Event, error) {
	var event structs.Event
	var events []structs.Event
	var err error
	rows, err := database.DB.Query("SELECT * FROM Event WHERE groupId = '" + groupId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var eventId, userId, groupId2, name, description, location, startTime, endTime, createdAt string
	for rows.Next() {
		rows.Scan(&eventId, &userId, &groupId2, &name, &description, &location, &startTime, &endTime, &createdAt)
		event = structs.Event{
			EventId:     eventId,
			UserId:      userId,
			GroupId:     groupId2,
			Name:        name,
			Description: description,
			Location:    location,
			StartTime:   startTime,
			EndTime:     endTime,
			CreatedAt:   createdAt,
		}
		events = append([]structs.Event{event}, events...)
	}
	return events, nil
}

func AllEventByUser(userId string, database *structs.DB) ([]structs.Event, error) {
	var group structs.Group
	var groups []structs.Group
	var err error
	rows, err := database.DB.Query("SELECT * FROM Groups ")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var groupId, admin, name, description, createdAt string
	for rows.Next() {
		rows.Scan(&groupId, &admin, &name, &description, &createdAt)
		group = structs.Group{
			CreatedAt:   createdAt,
			Name:        name,
			GroupID:     groupId,
			Description: description,
			Admin:       admin,
		}
		groups = append([]structs.Group{group}, groups...)
	}
	return groups, nil
}
