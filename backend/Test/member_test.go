package Test

import (
	"fmt"
	"testing"
	"time"

	"backend/pkg/auth"
	"backend/pkg/group"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

var (
	mnewUser  structs.User
	mnewUser2 structs.User
	mnewUser3 structs.User
)

func TestCreateMember(t *testing.T) {
	//----------- Setup -------------------

	mnewUser = CreateUser(database, t)
	mnewUser2 = CreateUser(database, t)
	mnewUser3 = CreateUser(database, t)
	userIdTest := mnewUser.UserId
	receiverId := mnewUser2.UserId

	group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: mnewUser2.UserId}
	groupId, err := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
	if err != nil {
		t.Errorf("Error Creating a Group into the db %v", err)
	}
	grouoIdTest := groupId

	createdAt := time.Now().Format("2006 January 02 3:4:5 pm")
	stmt, _ := database.DB.Prepare(`
		INSERT INTO GroupMember values (?, ?, ?)
	`)
	_, err3 := stmt.Exec(grouoIdTest, mnewUser3.UserId, createdAt)

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
		err := member.AcceptInvitationNotif(grouoIdTest, receiverId, userIdTest, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
	email1 := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email1, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	email2 := "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email2, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user2, *database)
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *database)
	groupId, err = group.CreateGroup("Hello", "Greating", result1.UserId, database)
	if err != nil {
		t.Errorf("Error Creating the group %v", err)
	}
	member.AddInvitationNotif(groupId, result1.UserId, result2.UserId, "invite", database)
	t.Run("Get all member notif of the user", func(t *testing.T) {
		notif, err := member.GetInvitationNotif(result2.UserId, database)
		l.LogMessage("TestGetMemberNotif", "GetMemberNotif", notif)
		if err != nil {
			t.Errorf("GetInvitationNotif Error %v", err)
		}
	})

	t.Run("Get all group invite notif", func(t *testing.T) {
		notif, err := member.GetAllInvitationNotif(database)
		l.LogMessage("TestGetMemberNotif", "GetMemberNotif", notif)
		if err != nil {
			t.Errorf("GetInvitationNotif Error %v", err)
		}
		if len(notif) == 0 {
			t.Errorf("Error - Declince %v", err)
		}

	})

	t.Run("Decline member notif", func(t *testing.T) {
		err := member.DeclineInvitationNotif(groupId, result2.UserId, result1.UserId, database)
		if err != nil {
			t.Errorf("Error - Declince %v", err)
		}
	})
}
