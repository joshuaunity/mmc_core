package schemas

import "time"

type UserRoomRole string

const (
	OWNER  UserRoomRole = "owner"
	ADMIN  UserRoomRole = "admin"
	MEMBER UserRoomRole = "member"
)

type Room struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"` // Type of room (group, private)
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserRoom struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"` // Role of the user in the room (owner, admin, member)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoomRecentActivity struct {
	Time   time.Time `json:"time"`
	RoomID string    `json:"room_id"`
}
