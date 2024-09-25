package models

import "time"

type Post struct {
	ID               int       `json:"id"`
	UserID           int       `json:"userId"`
	Content          string    `json:"content"`
	Category         string    `json:"category"`
	MalePercentage   int       `json:"malePercentage"`
	FemalePercentage int       `json:"femalePercentage"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
