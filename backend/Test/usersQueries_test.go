package Test

import (
	"testing"

	"backend/pkg/follow"
	"backend/pkg/users"
)

func TestGetAllPublicUsersNotFollowed(t *testing.T) {
	delete, _ := database.DB.Prepare(`DELETE FROM User`)
	delete.Exec()
	one := CreateUser(database, t)

	for i := 0; i < 3; i++ {
		temp := CreateUser(database, t)
		if i == 0 {
			follow.InsertFollow(one.UserId, temp.UserId, database)
		}
	}

	got, _ := users.GetAllPublicUsersNotFollowed(one.UserId, *database)

	if len(got) != 2 {
		t.Errorf("Len should be 2 but is %v", len(got))
	}
}
