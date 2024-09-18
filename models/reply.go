package models

import "time"

// Reply represents a reply to a post in the system
type Reply struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postId"`
	SenderID  int       `json:"senderId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
