package structs

type Follower struct {
	FollowingId string `json:"followingId"`
	FollowerId  string `json:"followerId"`
	CreatedAt   string `json:"createdAt"`
}

type FollowerNotif struct {
	UserId      string `json:"userId"`
	FollowingId string `json:"followingId"`
	CreatedAt   string `json:"createdAt"`
	Status      string `json:"status"`
	Read        int    `json:"read"`
}
