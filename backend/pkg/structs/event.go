package structs

type Event struct {
	EventId     string
	UserId      string `json:"UserId"`
	GroupId     string `json:"GroupId"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Location    string `json:"Location"`
	StartTime   string `json:"StartTime"`
	EndTime     string `json:"EndTime"`
	CreatedAt   string
	Status      bool `json:"Status"`
	ImageUpload string `json:"ImageUpload"`
}

type EventNotification struct {
	EventId string `json:"eventId"`
	UserId  string `json:"userId"`
	Read    int    `json:"read"`
}
