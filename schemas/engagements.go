package schemas

import "time"

type CommentSchema struct {
	ID         string    `json:"id"`
	User       string    `json:"user"`
	ResourceId string    `json:"resource_id"`
	Content    string    `json:"content"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type LikeSchema struct {
	User       string `json:"user"`
	ResourceId string `json:"resource_id"`
}

type FollowSchema struct {
	Follower  string `json:"follower"`
	Following string `json:"following"`
}
