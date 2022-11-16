package closefriend

import (
	"log"
	"time"

	"backend/pkg/helper"
	"backend/pkg/structs"
)

func UpdateCloseFriend(userId, closeFriendId string, database structs.DB) string {
	rows, err := database.DB.Query("SELECT * FROM CloseFriends WHERE closeFriendId = ? AND userId = ?", closeFriendId, userId)
	if err != nil {
		log.Println("Error Selecting the closefriends row: ", err)
		return "Error"
	}
	present := false
	defer rows.Close()
	for rows.Next() {
		present = true
	}
	if present { // Friendship exsists so remove
		if DeleteCloseFriend(userId, closeFriendId, database) {
			return "Removed"
		} else {
			return "Error"
		}
	}

	if AddCloseFriend(userId, closeFriendId, database) { // Add the friendship
		return "Added"
	} else {
		return "error"
	}
}

func DeleteCloseFriend(userId, closeFriendId string, database structs.DB) bool {
	qry, err := database.DB.Prepare("DELETE FROM CloseFriends WHERE userId = ? AND closeFriendId = ?")
	if err != nil {
		log.Println("Error preparing closefriends row: ", err)
		return false
	}
	_, execErr := qry.Exec(userId, closeFriendId)
	if execErr != nil {
		log.Println("Error executing closefriends row: ", execErr)
		return false
	}
	return true
}

func AddCloseFriend(userId, closeFriendId string, database structs.DB) bool {
	qry, err := database.DB.Prepare("INSERT INTO CloseFriends VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error Preparing closefriends insert row: ", err)
		return false
	}
	_, execErr := qry.Exec(userId, closeFriendId, time.Now().String())
	if execErr != nil {
		log.Println("Error executing closefriends insert row: ", err)
		return false
	}
	return true
}

func GetCloseFriends(userId string, database structs.DB) []structs.Info {
	rows, err := database.DB.Query("SELECT closeFriendId FROM CloseFriends WHERE userId = ?", userId)
	if err != nil {
		log.Println("Error Querying closefrineds table to get closefriends")
		return []structs.Info{}
	}
	var result []string
	var temp string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&temp)
		result = append(result, temp)
	}
	var closeFriends []structs.Info
	for _, v := range result {
		temp, getErr := helper.GetUserInfo(v, &database)
		if getErr != nil {
			log.Println("Error getting close friend information")
			return []structs.Info{}
		}
		closeFriends = append(closeFriends, temp)
	}
	return closeFriends
}
