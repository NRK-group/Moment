package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/handler"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)
var privacyValues = []int{0,1}
var wants = []string{`{"Message":"pending"}`,`{"Message":"follow"}`}
func TestFollowReq(t *testing.T) {
	t.Run("Pending request", func(t *testing.T) {
		// Insert a user that follows a new one
		for i, v := range privacyValues {

			emailOne := "follow@" + uuid.NewV4().String() + ".com"
			emailTwo := "follow@" + uuid.NewV4().String() + ".com"
			Env := handler.Env{Env: database}
	
	
			userOne := structs.User{FirstName: "Follow", LastName: "Follow", NickName: "Follow", AboutMe: "Follow", Email: emailOne, Password: "Password123", DateOfBirth: "06-08-2002"}
			userTwo := structs.User{FirstName: "FollowTwo", LastName: "FollowTwo", NickName: "FollowTwo", AboutMe: "FollowTwo", Email: emailTwo, Password: "Password123", DateOfBirth: "06-08-2002", IsPublic: v}
			auth.InsertUser(userOne, *database)
			auth.InsertUser(userTwo, *database)
			var testOne, testTwo structs.User
			auth.GetUser("email", emailOne, &testOne, *database)
			auth.GetUser("email", emailTwo, &testTwo, *database)
			// login userone
			loginOne := structs.User{Email: emailOne, Password: "Password123"}
			loginBytes, err := json.Marshal(loginOne)
			if err != nil {
				t.Errorf("Error marshalling the sampleUser")
				return
			}
	
			// Create the bytes into a reader
			testlogin := bytes.NewReader(loginBytes)
			req := httptest.NewRequest(http.MethodPost, "/login", testlogin)
			w := httptest.NewRecorder()
			Env.Login(w, req)
			//Now let user one follow user two
			followPending := structs.Follower{FollowerId: testOne.UserId, FollowingId: testTwo.UserId }
			followReqBytes, _ := json.Marshal(followPending)
			followReader := bytes.NewReader(followReqBytes)
	
			reqFol := httptest.NewRequest(http.MethodPut, "/followrequest", followReader)
			reqFol.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
			wr := httptest.NewRecorder()
			Env.FollowReq(wr, reqFol)
			got := wr.Body.String()
			t.Log("GOT === ", got)
			t.Log("WANT === ", wants[i])
			if got != wants[i] {
				t.Errorf("got %v \n want %v", got, wants[i])
			}
		}
	})
	t.Run("Unfollow the user", func(t *testing.T) {
		emailOne := "follow@" + uuid.NewV4().String() + ".com"
			emailTwo := "follow@" + uuid.NewV4().String() + ".com"
			Env := handler.Env{Env: database}
	
	
			userOne := structs.User{FirstName: "Follow", LastName: "Follow", NickName: "Follow", AboutMe: "Follow", Email: emailOne, Password: "Password123", DateOfBirth: "06-08-2002"}
			userTwo := structs.User{FirstName: "FollowTwo", LastName: "FollowTwo", NickName: "FollowTwo", AboutMe: "FollowTwo", Email: emailTwo, Password: "Password123", DateOfBirth: "06-08-2002", IsPublic: 1}
			auth.InsertUser(userOne, *database)
			auth.InsertUser(userTwo, *database)
			var testOne, testTwo structs.User
			auth.GetUser("email", emailOne, &testOne, *database)
			auth.GetUser("email", emailTwo, &testTwo, *database)
			// login userone
			loginOne := structs.User{Email: emailOne, Password: "Password123"}
			loginBytes, err := json.Marshal(loginOne)
			if err != nil {
				t.Errorf("Error marshalling the sampleUser")
				return
			}
	
			// Create the bytes into a reader
			testlogin := bytes.NewReader(loginBytes)
			req := httptest.NewRequest(http.MethodPost, "/login", testlogin)
			w := httptest.NewRecorder()
			Env.Login(w, req)
			//Now let user one follow user two
			followPending := structs.Follower{FollowerId: testOne.UserId, FollowingId: testTwo.UserId }
			followReqBytes, _ := json.Marshal(followPending)
			followReader := bytes.NewReader(followReqBytes)
			follow.FollowUser(testOne.UserId, testTwo.UserId, database)
	
			reqFol := httptest.NewRequest(http.MethodPut, "/followrequest", followReader)
			reqFol.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
			wr := httptest.NewRecorder()
			Env.FollowReq(wr, reqFol)
			got := wr.Body.String()
			want := `{"Message":"unfollow"}`
			if got != want {
				t.Errorf("got %v \n want %v", got, want)
			}
	})
}
