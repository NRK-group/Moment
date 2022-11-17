package Test

import (
	"testing"

	"backend/pkg/commets"
	"backend/pkg/post"
	"backend/pkg/structs"
)

func TestCreateCommets(t *testing.T) {
	t.Run("Insert comment to DB", func(t *testing.T) {
		content := "test23"
		newUser := CreateUser(database, t)
		post1 := structs.Post{UserID: newUser.UserId, Content: "hey", GroupID: "3233234", Image: "wasfdfgfd"}
		postId, err := post.CreatePost(post1.UserID, post1.Content, post1.GroupID, post1.Image, database)
		_, errc := commets.CreateComment(newUser.UserId, postId, content, database)
		if errc != nil {
			t.Errorf("Error inputing a comment in the db %v", err)
		}

		retrunComment, err := commets.GetComments(postId, database)

		if errc != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		if len(retrunComment) <= 0 {
			t.Errorf("the length of the comments array should be over 0 but it's %v ",
				len(retrunComment))
		}
	})
}