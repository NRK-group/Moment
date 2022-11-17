package structs

import "time"

type Follower struct {
	FollowingId string    `json:"followingId"`
	FollowerId  string    `json:"followerId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type FollowerNotif struct {
	UserId      Info      `json:"userId"`
	FollowingId Info      `json:"followingId"`
	CreatedAt   time.Time `json:"createdAt"`
	Status      string    `json:"status"`
	Read        int       `json:"read"`
}
