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
		SELECT id, user_id, content, category, male_percentage, female_percentage, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	row := config.DB.QueryRow(query, id)

	var post models.Post
	if err := row.Scan(
		&post.ID, &post.UserID, &post.Content, &post.Category,
		&post.MalePercentage, &post.FemalePercentage, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, fmt.Errorf("post not found")
		}
		return models.Post{}, fmt.Errorf("failed to scan post: %w", err)
	}
	return post, nil
}

func GetAllPosts() ([]models.Post, error) {
	query := `
		SELECT id, user_id, content, category, male_percentage, female_percentage, created_at, updated_at
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
		if err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.Category,
			&post.MalePercentage, &post.FemalePercentage, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CreatePost(post *dto.CreatePostRequest) (int, error) {
	query := `
		INSERT INTO posts (user_id, content, category, male_percentage, female_percentage, created_at, updated_at)
		VALUES ($1, $2, $3, 50, 50, NOW(), NOW())
		RETURNING id
	`

	var postID int
	err := config.DB.QueryRow(query, post.UserID, post.Content, post.Category).Scan(&postID)
	if err != nil {
		return 0, fmt.Errorf("failed to create post: %w", err)
	}
	return postID, nil
}

func UpdatePostPercentage(postID string, malePercentage int, femalePercentage int) error {
	query := `
		UPDATE posts
		SET male_percentage = $1, female_percentage = $2, updated_at = NOW()
		WHERE id = $3
	`

	_, err := config.DB.Exec(query, malePercentage, femalePercentage, postID)
	if err != nil {
		return fmt.Errorf("failed to update percentages: %w", err)
	}
	return nil
}
