package services

import (
	"database/sql"
	"fmt"
	"log"
	"morpet-backend/config"
	"morpet-backend/handlers/dto"
	"morpet-backend/models"
)

func GetUserById(id string) (models.User, error) {
	query := `
		SELECT id, name, age, profile, profile_image_url, gender, mbti, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	// 単一の結果を取得するために QueryRow を使用
	row := config.DB.QueryRow(query, id)

	var user models.User
	// 結果を user 構造体にマッピング
	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Profile, &user.ProfileImageUrl, &user.Gender, &user.Mbti, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, fmt.Errorf("failed to scan user: %w", err)
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	query := `
		SELECT id, name, age, profile, profile_image_url, gender, mbti, created_at, updated_at
		FROM users
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Profile, &user.ProfileImageUrl, &user.Gender, &user.Mbti, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(req *dto.CreateUserRequest) (int64, error) {
	query := `
		INSERT INTO users (name, age, profile, profile_image_url, gender, mbti, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id
	`

	var userID int64
	// クエリを実行し、挿入されたユーザーのIDを取得
	err := config.DB.QueryRow(query, req.Name, req.Age, req.Profile, req.ProfileImageUrl, req.Gender, req.Mbti).Scan(&userID)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return userID, nil
}

func UpdateUser(id string, req *dto.UpdateUserRequest) error {
	query := `
		UPDATE users 
        SET age = $1, profile = $2, profile_image_url = $3, gender = $4, mbti = $5, updated_at = NOW()
        WHERE id = $6
	`
	_, err := config.DB.Exec(query, req.Age, req.Profile, req.ProfileImageUrl, req.Gender, req.Mbti, id)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
