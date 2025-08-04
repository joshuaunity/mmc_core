package schemas

import "time"

type Comment struct {
	ID         string    `json:"id"`
	User       string    `json:"user"`
	ResourceId string    `json:"resource_id"`
	Content    string    `json:"content"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Like struct {
	ID         string    `json:"id"`
	User       string    `json:"user"`
	ResourceId string    `json:"resource_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Follow struct {
	Follower  string `json:"follower"`
	Following string `json:"following"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
