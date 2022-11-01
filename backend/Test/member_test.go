package Test

import (
	"fmt"
	"testing"
	"time"

	"backend/pkg/member"
)

func TestCreateMember(t *testing.T) {

	//----------- Setup -------------------
	grouoIdTest := "dhgfhfdj"
	userIdTest := "esfesfes5f"+time.Now().String()
	database := DatabaseSetup()
	createdAt := time.Now().Format("2006 January 02 3:4:5 pm")
	stmt, _ := database.DB.Prepare(`
		INSERT INTO GroupMember values (?, ?, ?)
	`)
	_, err3:= stmt.Exec(grouoIdTest, "userIdTest", createdAt)

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
		Member, err := member.GetMembers("grouoIdTest465", database)
		if err != nil || len(Member) > 0 {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Add member to an unknow group", func(t *testing.T) {
		str, err := member.AddMember("grouoIdTest", userIdTest, database)
		fmt.Println(str)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("Add member to the same group", func(t *testing.T) {
		str, err := member.AddMember(grouoIdTest, userIdTest, database)
		fmt.Println(err)

		want := false
		got := str

		if got != want {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}
