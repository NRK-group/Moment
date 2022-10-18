package queries

import (
	"database/sql"
	"fmt"
	"backend/pkg/handler"

	uuid "github.com/satori/go.uuid"
)

type DB struct {
	DB *sql.DB
}

func (DB DB) InsertUser(newUser handler.User) error {
	userId := uuid.NewV4().String()

	stmt, err := DB.DB.Prepare(`INSERT INTO User values`)
	if err != nil {
		fmt.Println("Error preparing inserting user into the db: ", err)
		return err
	}

	_, execErr := stmt.Exec(userId, "", newUser.FirstName, newUser.LastName, newUser.NickName, newUser.Email, newUser.DateOfBirth, newUser.Avatar, newUser.AboutMe, newUser.CreatedAt, newUser.IsLoggedIn, newUser.IsPublic, newUser.NumFollowers, newUser.NumFollowing, newUser.NumPosts, newUser.Password)

	if execErr != nil {
		fmt.Println("Error executing inserting user into the db: ", execErr)
		return execErr
	}
	return nil
}
