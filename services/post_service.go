package services

import (
	"database/sql"
	"fmt"
	"morpet-backend/config"
	"morpet-backend/models"
)

func GetPostbyId(id string) (models.Post, error) {
	query := `
		SELECT id, user_id, content, category, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	row := config.DB.QueryRow(query, id)

	var post models.Post
	if err := row.Scan(
		&post.ID, &post.UserID, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, fmt.Errorf("reply not post")
		}
		return models.Post{}, fmt.Errorf("failed to scan post: %w", err)
	}
	return post, nil
}
