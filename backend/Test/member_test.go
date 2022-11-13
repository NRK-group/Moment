package Test

import (
	"fmt"
	"testing"
	"time"

	l "backend/pkg/log"
	"backend/pkg/member"
)

func TestCreateMember(t *testing.T) {
	//----------- Setup -------------------
	grouoIdTest := "dhgfhfdj"
	userIdTest := "esfesfesf"
	receiverId := "hello"
	createdAt := time.Now().Format("2006 January 02 3:4:5 pm")
	stmt, _ := database.DB.Prepare(`
		INSERT INTO GroupMember values (?, ?, ?)
	`)
	_, err3 := stmt.Exec(grouoIdTest, "userIdTest", createdAt)

	if err3 != nil {
		fmt.Println("inside AddMember", err3)
		return
	}

	t.Run("Add member to group", func(t *testing.T) {
		str, err := member.AddMember(grouoIdTest, userIdTest, database)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Get all members from a group", func(t *testing.T) {
		Member, err := member.GetMembers(grouoIdTest, database)
		fmt.Println(Member)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("test wrong unknow group ID", func(t *testing.T) {
		Member, _ := member.GetMembers("grouoIdTest465", database)
		if len(Member) > 0 {
			t.Errorf("Error got %v want %v", len(Member), 0)
		}
	})
	t.Run("Add member to an unknow group", func(t *testing.T) {
		str, err := member.AddMember("grouoIdTest", userIdTest, database)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	t.Run("Add invite notif", func(t *testing.T) {
		err := member.AddInvitationNotif(grouoIdTest, userIdTest, receiverId, "invite", database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	t.Run("Accept invite notif", func(t *testing.T) {
		err := member.AcceptInvitationNotif(grouoIdTest, userIdTest, receiverId, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	member.AddInvitationNotif(grouoIdTest, "Hello1", "Member1", "invite", database)
	member.AddInvitationNotif(grouoIdTest, "Hello2", "Member1", "invite", database)
	t.Run("Get all member notif of the user", func(t *testing.T) {
		notif, err := member.GetInvitationNotif("Member1", database)
		l.LogMessage("TestGetMemberNotif", "GetMemberNotif", notif)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	t.Run("Decline member notif", func(t *testing.T) {
		err := member.DeclineInvitationNotif(grouoIdTest, userIdTest, receiverId, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}
