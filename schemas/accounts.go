package schemas

import (
	"time"
)

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Handle     string    `json:"handle"`
	PhotoURL   string    `json:"photo_url"`
	Groups     []Group   `json:"groups"`
	IsArchived bool      `json:"is_archived"`
	IsActive   bool      `json:"is_active"`
	IsVerified bool      `json:"is_verified"`
	LastLogin  time.Time `json:"last_login"`
	CreatedAt  time.Time `json:"created_at"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StreamingServiceToken struct {
	ID           string    `gorm:"type:string;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID       string    `json:"user_id"`
	Service      string    `json:"service"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}