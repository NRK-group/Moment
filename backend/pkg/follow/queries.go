package follow

import (
	"time"

	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/structs"
)

// InsertFollow will insert a follow into the database
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	database: the database to insert the follow into
func InsertFollow(follwerId, followingId string, database *structs.DB) error {
	createdAt := time.Now().String()
	stmt, err := database.DB.Prepare("INSERT INTO Follower (followerId, followingId, createdAt) VALUES (?, ?, ?)")
	if err != nil {
		l.LogMessage("follow.go", "InsertFollow", err)
		return err
	}
	_, err = stmt.Exec(follwerId, followingId, createdAt)
	if err != nil {
		l.LogMessage("follow.go", "InsertFollow", err)
		return err
	}
	return nil
}

// DeleteFollow will delete a follow from the database
//
// Params:
//
//	followerId: the id of the user who is following
//	followingId: the id of the user who is being followed
//	database: the database to delete the follow from
func DeleteFollow(followerId, followingId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("DELETE FROM Follower WHERE followerId = ? AND followingId = ?")
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	_, err = stmt.Exec(followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	stmt, err = database.DB.Prepare("DELETE FROM FollowNotif WHERE userId = ? AND followingId = ?")
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	_, err = stmt.Exec(followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	return nil
}

// CheckIfFollow checks if current user is following the other user
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	database: the database to check the follow in
func CheckIfFollow(followerId, followingId string, database *structs.DB) bool {
	// check if a user is following another user
	stmt, err := database.DB.Query("SELECT * FROM Follower WHERE followerId = ? AND followingId = ?", followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "CheckFollow", err)
		return false
	}
	var follower structs.Follower
	for stmt.Next() {
		stmt.Scan(
			&follower.FollowerId,
			&follower.FollowingId,
			&follower.CreatedAt,
		)
	}
	return follower != structs.Follower{}
}

// InsertFollowNotif will insert a follow notification into the database
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	status: the status of the follow (follow or pending)
//	database: the database to insert the follow notification into
func InsertFollowNotif(followerId, followingId, status string, database *structs.DB) error {
	createdAt := time.Now().String()
	stmt, err := database.DB.Prepare("INSERT INTO FollowNotif (userId, followingId, status, createdAt, unread) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		l.LogMessage("follow.go", "InsertFollowNotif", err)
		return err
	}
	_, err = stmt.Exec(followerId, followingId, status, createdAt, true)
	if err != nil {
		l.LogMessage("follow.go", "InsertFollowNotif", err)
		return err
	}
	return nil
}

// FollowUser will follow a user if the user does not already follow the other user
// and delete the follow if the user does follow the other user
//
//	returns status (follow, pending or  unfollow) and error
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	database: the database to insert the follow into
func FollowUser(followerId, followingId string, database *structs.DB) (string, error) {
	if CheckIfFollow(followerId, followingId, database) {
		DeleteFollow(followerId, followingId, database)
		return "unfollow", nil
	}
	if helper.CheckUserIfPublic(followingId, database) {
		InsertFollow(followerId, followingId, database)
		InsertFollowNotif(followerId, followingId, "follow", database)
		return "follow", nil
	}
	InsertFollowNotif(followerId, followingId, "pending", database)
	return "pending", nil
}
