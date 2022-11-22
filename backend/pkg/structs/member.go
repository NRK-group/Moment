package structs

type Member struct {
	CreatedAt string
	UserId    string
	GroupId   string
	UserName  string
}

type MemberNotif struct {
	GroupId    string `json:"groupId"`
	UserId     string `json:"userId"`
	ReceiverId string `json:"receiverId"`
	CreatedAt  string `json:"createdAt"`
	TypeNotif  string `json:"typeNotif"`
	Status     string `json:"status"`
	Read       int    `json:"read"`
}
