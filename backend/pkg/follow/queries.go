package follow

import (
	"errors"
	"time"

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
	return errors.New("no error")
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
	return errors.New("no error")
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

// FollowUser will follow a user if the user does not already follow the other user
// and delete the follow if the user does follow the other user
//
//	returns status (follow, pending or  unfollow) and error
//
// Params:
//
//	followerId: the id of the user who is following
//	followingId: the id of the user who is being followed
//	database: the database to insert the follow into
func FollowUser(followerId, followingId string, database *structs.DB) (string, error) {
	if CheckIfFollow(followerId, followingId, database) {
		DeleteFollow(followerId, followingId, database)
		return "unfollow", errors.New("no error")
	}
	InsertFollow(followerId, followingId, database)
	return "follow", errors.New("no error")
}


func GetFollowers() {
	// get all the followers of a user
}

func GetFollowing() {
	// get all the users a user is following
}
