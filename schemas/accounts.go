package schemas

import (
	"time"
)


type UserSchema struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Handle     string    `json:"handle"`
	PhotoURL   string    `json:"photo_url"`
	IsArchived bool      `json:"is_archived"`
	IsActive   bool      `json:"is_active"`
	IsVerified bool      `json:"is_verified"`
	LastLogin  time.Time `json:"last_login"`
	CreatedAt  time.Time `json:"created_at"`
}
