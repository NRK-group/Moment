package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/pkg/event"
	"backend/pkg/group"
	"backend/pkg/handler"
	"backend/pkg/member"
	"backend/pkg/structs"
)

var database = DatabaseSetup()

func LoginUser(database *structs.DB, t *testing.T) (*httptest.ResponseRecorder, *handler.Env, structs.User) {
	newUser := CreateUser(database, t)
	loginStruct := structs.User{Email: newUser.Email, Password: "Password123"}

	loginUserBytes, err := json.Marshal(loginStruct)
	if err != nil {
		t.Errorf("Error marshalling the sampleUser")
	}
	Env := &handler.Env{Env: database}

	// Create the bytes into a reader
	loginReq := bytes.NewReader(loginUserBytes)
	req := httptest.NewRequest(http.MethodPost, "/login", loginReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)

	return w, Env, newUser
}

func TestHealthCheckGroupHandlerHttpGet(t *testing.T) {
	w, Env, _ := LoginUser(database, t)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqq, err := http.NewRequest("GET", "/group", nil)
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Env.Group)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, reqq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected2 := http.StatusOK
	if rr.Code != expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected2)
	}
}

func TestHealthCheckGroupHttpGet(t *testing.T) {
	w, Env, _ := LoginUser(database, t)

	reqq := httptest.NewRequest(http.MethodGet, "/group", nil)
	reqq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}

	Env.Group(w, reqq)
	want := 200
	got := w.Code

	if got != want {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestCreateGroup(t *testing.T) {
	t.Run("Creating a group", func(t *testing.T) {
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: newUser.UserId}
		_, err := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})

	t.Run("get all groups", func(t *testing.T) {
		_, err := group.AllGroups("6t78t8t87", database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
	})
}

func TestGroupHandlerGettingAllGroups(t *testing.T) {
	req, err := http.NewRequest("GET", "/group", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Env := &handler.Env{Env: database}
	handler := http.HandlerFunc(Env.Group)
	handler.ServeHTTP(rr, req)
	expected := http.StatusOK
	if status := rr.Code; status != expected && strings.Contains(rr.Body.String(), "GroupID") {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

var (
	groupID   string
	postIdArr []string
)

func TestGettingAllPostFromAGroup(t *testing.T) {
	t.Run("Creating 10 posts for a group", func(t *testing.T) {
		newUser := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: newUser.UserId}
		Id, err := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}
		groupID = Id
		var postId string

		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				postId = CreatePost("id2", database, t)
			} else {
				postId = CreatePost(groupID, database, t)
			}
			if i%2 != 0 {
				postIdArr = append([]string{postId}, postIdArr...)
			}
		}
	})

	t.Run("get all group posts", func(t *testing.T) {
		groups, err := group.AllGroupPosts(groupID, database)
		if err != nil {
			t.Errorf("Error Inserting the struct into the db %v", err)
		}

		var num int
		expected := 5

		for i := 0; i < len(groups); i++ {
			for r := 0; r < len(postIdArr); r++ {
				if postIdArr[r] == groups[i].PostID {
					num = num + 1
				}
			}
		}

		if num != expected {
			t.Errorf("Error posts doesn't match %d ", num)
		}
	})
}

func TestReadGroupNotif(t *testing.T) {
	t.Run("Read group notif", func(t *testing.T) {
		newUser := CreateUser(database, t)
		newUser2 := CreateUser(database, t)
		newUser3 := CreateUser(database, t)
		group1 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: newUser.UserId}
		Id, err := group.CreateGroup(group1.Name, group1.Description, group1.Admin, database)
		if err != nil {
			t.Errorf("Error creating group%v", err)
		}
		member.AddMember(Id, newUser2.UserId, database)
		member.AddMember(Id, newUser3.UserId, database)
		event1 := structs.Event{Name: "Test4", Description: "Eating Pie", Location: " Location ", StartTime: "StartTime", EndTime: " EndTime", UserId: newUser.UserId, GroupId: Id}
		_, err = event.AddEvent(Id, event1, database)
		if err != nil {
			t.Errorf("Error Inserting event %v", err)
		}
		group2 := structs.Group{Name: "Pie", Description: "Eating Pie", Admin: newUser.UserId}
		Id2, err := group.CreateGroup(group2.Name, group2.Description, group2.Admin, database)
		if err != nil {
			t.Errorf("Error creating group%v", err)
		}

		member.AddInvitationNotif(Id2, newUser.UserId, newUser2.UserId, "invite", database)
		err = group.ReadGroupNotif(newUser2.UserId, database)
		if err != nil {
			t.Errorf("Error reading group notif %v", err)
		}
	})
}
