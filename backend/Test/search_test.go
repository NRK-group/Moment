package Test

import (
	"testing"

	q "backend/pkg/search"
)

func TestSearchUsers(t *testing.T) {
	id := "hello"
	users, err := q.GetAllUsers(id, database)
	t.Run("GetAllUsers", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
		if len(users) == 0 {
			t.Error("No users found")
		}
	})
	// query := "Founde"
	// users, err = q.SearchUsers(id, query, database)
	// t.Run("SearchUsers", func(t *testing.T) {
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	if len(users) == 0 {
	// 		t.Error("No users found")
	// 	}
	// })
}
