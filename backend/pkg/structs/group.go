package structs

type Group struct {
	
	CreatedAt   string
	Name        string `json:"Name"`
	GroupID     string 
	Description string `json:"Description"`
	Admin       string `json:"Admin"`

}
