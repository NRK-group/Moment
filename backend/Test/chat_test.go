package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/follow"
	"backend/pkg/group"
	"backend/pkg/handler"
	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/member"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestPrivateChat(t *testing.T) {
	Env := handler.Env{Env: DatabaseSetup()}
	defer Env.Env.DB.Close()
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user1, *Env.Env)
	var result1 structs.User
	err := auth.GetUser("email", user1.Email, &result1, *Env.Env)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 1, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	err = auth.InsertUser(*user2, *Env.Env)
	t.Run("Insert user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error inserting test struct %v", err)
		}
	})
	var result2 structs.User
	err = auth.GetUser("email", user2.Email, &result2, *Env.Env)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	_, err = follow.FollowUser(result1.UserId, result2.UserId, Env.Env)
	t.Run("Follow user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error following user %v", err)
		}
	})
	_, err = follow.FollowUser(result2.UserId, result1.UserId, Env.Env)
	t.Run("Follow user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error following user %v", err)
		}
	})
	_, err = chat.InsertNewChat(result1.UserId, result2.UserId, Env.Env)
	t.Run("Insert new chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("InsertNewChat() error = %v", err)
		}
	})
	chats, err := chat.GetPreviousPrivateChat(result1.UserId, Env.Env)
	t.Run("Get Previous Private chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("GetPreviousPrivateChat() error = %v", err)
		}
	})
	t.Run("Get Previous Private chat", func(t *testing.T) {
		if len(chats) != 0 {
			t.Errorf("Error expected 0 chat, got %v", len(chats))
		}
	})
	err = chat.InsertNewGroupChat("hello", Env.Env)
	t.Run("Insert New Group Chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("InsertNewGroupChat() error = %v", err)
		}
	})
	_, err = helper.GetUserInfo(result1.UserId, Env.Env)
	t.Run("Get userinfo for the chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("GetUserInfoForChat() error = %v", err)
		}
	})
	b, r, err := chat.CheckIfChatExists(result1.UserId, result2.UserId, Env.Env)
	t.Run("Check if chat exists", func(t *testing.T) {
		if err != nil {
			t.Errorf("CheckIfChatExists() error = %v", err)
		}
	})
	t.Run("Check if chat exists", func(t *testing.T) {
		if b != true {
			t.Errorf("Error expected true, got %v", err)
		}
	})
	t.Run("Check if chat exists", func(t *testing.T) {
		if r.Details.Id != result2.UserId {
			t.Errorf("Error expected %v, got %v", result2.UserId, r)
		}
	})
	b, r, err = chat.CheckIfChatExists(result2.UserId, result1.UserId, Env.Env)
	t.Run("Check if chat exists", func(t *testing.T) {
		if err != nil {
			t.Errorf("CheckIfChatExists() error = %v", err)
		}
	})
	t.Run("Check if chat exists", func(t *testing.T) {
		if b != true {
			t.Errorf("Error expected true, got %v", err)
		}
	})
	t.Run("Check if chat exists", func(t *testing.T) {
		if r.Details.Id != result1.UserId {
			t.Errorf("Error expected %v, got %v", result1.UserId, r)
		}
	})
}

func TestGroupChat(t *testing.T) {
	Env := handler.Env{Env: DatabaseSetup()}
	defer Env.Env.DB.Close()
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}

	auth.InsertUser(*user1, *Env.Env)
	var result1 structs.User
	err := auth.GetUser("email", user1.Email, &result1, *Env.Env)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	err = auth.InsertUser(*user2, *Env.Env)
	t.Run("Insert user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error inserting test struct %v", err)
		}
	})
	var result2 structs.User
	err = auth.GetUser("email", user2.Email, &result2, *Env.Env)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	_, err = follow.FollowUser(result1.UserId, result2.UserId, Env.Env)
	t.Run("Follow user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error following user %v", err)
		}
	})
	_, err = follow.FollowUser(result2.UserId, result1.UserId, Env.Env)
	t.Run("Follow user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error following user %v", err)
		}
	})
	groupId, err := group.CreateGroup("hello", "testing", result1.UserId, Env.Env)
	fmt.Println(groupId)
	t.Run("Create group", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error creating group %v", err)
		}
	})
	_, err = member.AddMember(groupId, result2.UserId, Env.Env)
	t.Run("Add member", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error adding member %v", err)
		}
	})
	_, err = chat.GetUserGroups(result1.UserId, Env.Env)
	t.Run("Get user groups", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting user groups %v", err)
		}
	})
	_, err = chat.GetAllMembersOfGroup(groupId, Env.Env)
	t.Run("Get all members of group", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting all members of group %v", err)
		}
	})
	chatlist, err := chat.GetPreviousGroupChat(result1.UserId, Env.Env)
	l.LogMessage("Test", "GetPreviousGroupChat", chatlist)
	t.Run("Get previous group chat", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting previous group chat %v", err)
		}
	})
}

func TestChatHandler(t *testing.T) {
	Env := handler.Env{Env: DatabaseSetup()}
	defer Env.Env.DB.Close()
	validationValidEmail := "add" + uuid.NewV4().String() + "@test.com"
	inputUser := &structs.User{
		FirstName: "Adriell", LastName: "Validation123", NickName: "", Email: validationValidEmail, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "-", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 0, NumFollowing: 0, NumPosts: 0,
	}
	err := auth.InsertUser(*inputUser, *Env.Env)
	t.Run("Insert new user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error inserting test struct %v", err)
		}
	})
	var result structs.User
	err = auth.GetUser("email", validationValidEmail, &result, *Env.Env)
	t.Run("Get user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error getting test struct %v", err)
		}
	})
	sampleUser := &structs.User{
		Email: validationValidEmail, Password: "Password123",
	}
	sampleUserBytes, err := json.Marshal(sampleUser)
	t.Run("Marshal user", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
	})
	testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
	req := httptest.NewRequest(http.MethodPost, "/login", testReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	t.Run("Login user", func(t *testing.T) {
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	nr := httptest.NewRequest(http.MethodGet, "/chat", nil) // create a request
	nr.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	nrr := httptest.NewRecorder() // create a response recorder
	Env.Chat(nrr, nr)             // call the handler
	// get cookies
	cookie, _ := nr.Cookie("session_token")
	got, _ := auth.SliceCookie(cookie.Value)
	want := result.UserId
	t.Run("Get cookie", func(t *testing.T) {
		if got[0] != want {
			t.Errorf("Expected %s, got %s", want, got[0])
		}
	})
	t.Run("Get chat", func(t *testing.T) {
		if nrr.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, nrr.Code)
		}
	})
	type Receiver struct {
		Id string `json:"receiverId"`
	}

	receiver := &Receiver{
		Id: result.UserId,
	}
	receiverBytes, err := json.Marshal(receiver)
	rBody := bytes.NewReader(receiverBytes)
	nr2 := httptest.NewRequest(http.MethodPost, "/chat", rBody) // create a request
	nr2.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	nrr2 := httptest.NewRecorder() // create a response recorder
	Env.Chat(nrr2, nr2)            // call the handler
	// get cookies
	cookie2, _ := nr.Cookie("session_token")
	got2, _ := auth.SliceCookie(cookie2.Value)
	want2 := result.UserId
	t.Run("Get cookie", func(t *testing.T) {
		if got2[0] != want2 {
			t.Errorf("Expected %s, got %s", want2, got2[0])
		}
	})
	t.Run("Get chat", func(t *testing.T) {
		if nrr2.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, nrr2.Code)
		}
	})
}
