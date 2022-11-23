package users

import (
	"log"

	"backend/pkg/structs"
)

func GetAllPublicUsersNotFollowed(userId string, db *structs.DB) ([]structs.Info, error) {
	following, qryErr := db.DB.Query(`SELECT followingId FROM Follower WHERE followerId = ?`, userId)
	if qryErr != nil {
		log.Println("Error getting following users: ", qryErr)
		return []structs.Info{}, qryErr
	}
	var follow []string
	defer following.Close()
	for following.Next() {
		var temp string
		following.Scan(&temp)
		follow = append(follow, temp)
	}
	rows, err := db.DB.Query(`SELECT userId, firstName, lastName, nickName, avatar FROM User WHERE isPublic = 1`)
	if err != nil {
		log.Println("Error getting public non followed users: ", err)
		return []structs.Info{}, err
	}
	defer rows.Close()
	var result []structs.Info
	for rows.Next() {
		var temp structs.Info
		rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Name, &temp.Img)
		contains := false
		for _, v := range follow {
			if temp.Id == v || temp.Id == userId {
				contains = true
				break
			}
		}
		if temp.Name == "" {
			temp.Name = temp.FirstName + " " + temp.LastName
		}
		if !contains {
			result = append(result, temp)
		}
	}
	return result, nil
}
