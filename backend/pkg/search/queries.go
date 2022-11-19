package search

import "backend/pkg/structs"

// GeAllUsers returns all users
//
// Param:
//
// userId: user id
// database: the database
func GetAllUsers(userId string, database *structs.DB) ([]structs.Info, error) {
	var users []structs.Info
	var user structs.Info
	stmt, err := database.DB.Query("SELECT userId, firstName, lastName, nickName, avatar FROM User WHERE userId != ?", userId)
	if err != nil {
		return nil, err
	}
	for stmt.Next() {
		err = stmt.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Name, &user.Img)
		if err != nil {
			return nil, err
		}
		if user.Name == "" {
			user.Name = user.FirstName + " " + user.LastName
		}
		users = append([]structs.Info{user}, users...)
	}
	return users, nil
}