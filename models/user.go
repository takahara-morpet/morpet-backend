package models

import "time"

// dto.User
type User struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Age             int       `json:"age"`
	Profile         string    `json:"profile"`
	ProfileImageUrl string    `json:"profileImageUrl"`
	Gender          string    `json:"gender"`
	Mbti            string    `json:"mbti"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
