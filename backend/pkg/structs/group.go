package structs

type Group struct {
	CreatedAt   string   `json:"CreatedAt"`
	Name        string   `json:"Name"`
	GroupID     string   `json:"GroupID"`
	Img         string   `json:"Img"`
	Description string   `json:"Description"`
	Admin       string   `json:"Admin"`
	Members     []Member `json:"Members"`
	Member      bool   `json:"member"`
}
