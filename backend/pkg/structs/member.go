package structs

type Member struct {
	CreatedAt string
	UserId    string
	GroupId   string
}

type MemberNotif struct {
	GroupId    string
	UserId     string
	ReceiverId string
	CreatedAt  string
	TypeNotif  string
	Status     string
	Read       int
}
