package structs

type Event struct {
	EventId     string
	UserId      string
	GroupId     string
	Name        string
	Description string
	Location    string
	StartTime   string
	EndTime     string
	CreatedAt   string
}

type EventNotification struct {
	EventId string `json:"eventId"`
	UserId  string `json:"userId"`
	Read    int    `json:"read"`
}
