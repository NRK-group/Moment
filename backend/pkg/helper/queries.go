package helper

import (
	"backend/pkg/structs"
)

// CheckUserStatus
func CheckUserIfPublic(userId string, database *structs.DB) bool {
	stmt, _ := database.DB.Query("SELECT isPublic FROM User WHERE userId = ?", userId)
	var status int
	for stmt.Next() {
		stmt.Scan(&status)
	}
	return status == 1
}
