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

// GetUserInfo returns the user info for the chat writer
//
// Param:
//
//	userId: the user id
//	database: the database
func GetUserInfo(userId string, database *structs.DB) (structs.Info, error) {
	var userInfo structs.Info
	var user structs.User
	stmt, err := database.DB.Query("SELECT userId, firstName, lastName, nickName, avatar FROM User WHERE userId = ?", userId)
	if err != nil {
		return userInfo, err
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar)
		if err != nil {
			return structs.Info{}, err
		}
		userInfo = structs.Info{
			Id:        user.UserId,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Img:       user.Avatar,
		}
		if user.NickName != "" {
			userInfo.Name = user.NickName
		} else {
			userInfo.Name = user.FirstName + " " + user.LastName
		}
	}
	return userInfo, nil
}
