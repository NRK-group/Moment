package helper

import (
	"backend/pkg/structs"
)

// CheckUserStatus
func CheckUserPrivacy(userId string, database *structs.DB) string {
	stmt, err := database.DB.Query("SELECT isPublic FROM User WHERE userId = ?", userId)
	if err != nil {
		return "error"
	}
	var status int
	for stmt.Next() {
		stmt.Scan(&status)
	}
	if status == 0 {
		return "private"
	}
	return "public"
}
