package structs

import "time"

type Event struct {
	EventId           string    `json:"EventId"`
	UserId            string    `json:"UserId"`
	GroupId           string    `json:"GroupId"`
	Name              string    `json:"Name"`
	Description       string    `json:"Description"`
	Location          string    `json:"Location"`
	StartTime         string    `json:"StartTime"`
	EndTime           string    `json:"EndTime"`
	CreatedAt         time.Time `json:"CreatedAt"`
	Status            string    `json:"Status"`
	ImageUpload       string    `json:"ImageUpload"`
	NumOfParticipants int
	UserName          string
	Participants      []EventParticipant
}

type EventNotification struct {
	EventId string `json:"eventId"`
	UserId  string `json:"userId"`
	Read    int    `json:"read"`
}
