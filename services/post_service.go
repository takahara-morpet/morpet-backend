package services

import (
	"database/sql"
	"fmt"
	"morpet-backend/config"
	"morpet-backend/handlers/dto"
	"morpet-backend/models"
)

func GetPostById(id string) (models.Post, error) {
	query := `
		SELECT id, user_id, content, category, created_at, updated_at
		FROM posts
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
func GetAllPosts() ([]models.Post, error) {
	query := `
		SELECT id, user_id, content, category, created_at, updated_at
		FROM posts
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve posts: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func CreatePost(post *dto.CreatePostRequest) (int, error) {
	query := `
		INSERT INTO posts (user_id, content, category, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`

	var postID int
	err := config.DB.QueryRow(query, post.UserID, post.Content, post.Category).Scan(&postID)
	if err != nil {
		return 0, fmt.Errorf("failed to create post: %w", err)
	}
	return postID, nil
}
