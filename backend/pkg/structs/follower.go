package structs

type Follower struct {
	FollowingId string `json:"followingId"`
	FollowerId  string `json:"followerId"`
	CreatedAt   string `json:"createdAt"`
}
