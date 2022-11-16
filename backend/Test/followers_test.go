package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/follow"
	"backend/pkg/handler"
	"backend/pkg/structs"
)

func TestFollowers(t *testing.T) {
	Env := handler.Env{Env: database}

	userOne := CreateUser(database, t)
	userTwo := CreateUser(database, t)
	userThree := CreateUser(database, t)
	follow.FollowUser(userTwo.UserId, userOne.UserId, database)
	follow.FollowUser(userThree.UserId, userOne.UserId, database)
	// Log user one in
	loginStruct := structs.User{Email: userOne.Email, Password: userOne.Password}
	logUser, _ := json.Marshal(loginStruct)
	// Create the bytes into a reader
	logReq := bytes.NewReader(logUser)
	req := httptest.NewRequest(http.MethodPost, "/login", logReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	followReq := httptest.NewRequest(http.MethodGet, "/followers", nil)
	followReq.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	wr := httptest.NewRecorder()
	Env.Followers(wr, followReq)
	var got []structs.Info
	json.Unmarshal(wr.Body.Bytes(), &got)
	for _, v := range got {
		if v.Id != userTwo.UserId || v.Id != userThree.UserId {
			t.Errorf("got %v, want: %v OR %v", v, userThree.UserId, userTwo.UserId)
		}
	}
}
