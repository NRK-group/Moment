package Test

import (
	"testing"

	"backend/pkg/auth"
	"backend/pkg/commets"
	"backend/pkg/post"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestCreateCommets(t *testing.T) {
	t.Run("Insert comment to DB", func(t *testing.T) {
		content := "test23"
		post1 := structs.Post{UserID: "3232131221", Content: "hey", GroupID: "3233234", Image: "wasfdfgfd"}
		postId, err := post.CreatePost(post1.UserID, post1.Content, post1.GroupID, post1.Image, database)

		currentEmail := "hello" + uuid.NewV4().String() + "@test.com"
		currentUser := &structs.User{
			FirstName: "FirstTest", LastName: "LastTest", NickName: "NickTest", Email: currentEmail, Password: "Password123",
			DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
			IsLoggedIn: 0, IsPublic: 1, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
		}
		auth.InsertUser(*currentUser, *database)
		var currentResult structs.User
		auth.GetUser("email", currentEmail, &currentResult, *database)

		_, errc := commets.CreateComment(currentResult.UserId, postId, content, "", database)
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
