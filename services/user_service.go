package services

import (
	"fmt"
	"log"
	"morpet-backend/config"
	"morpet-backend/handlers/dto"
	"morpet-backend/models"
)

func GetAllUsers() ([]models.User, error) {
	query := "SELECT id, name, age FROM users"

	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Failed to query", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Println("Failed to scan a row", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("An error occured on setting results", err)
		return nil, err
	}

	return users, nil
}

func CreateUser(req *dto.CreateUserRequest) error {
	query := `
		INSERT INTO users (name, age, created_at,updated_at) 
		VALUES ($1, $2, NOW(),NOW())
	`

	// クエリを実行し、エラーハンドリング
	_, err := config.DB.Exec(query, req.Name, req.Age)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
func UpdateUser(id string, req *dto.UpdateUserRequest) error {
	query := `
		UPDATE users 
        SET age = $1
        WHERE id = $2
	`
	_, err := config.DB.Exec(query, req.Age, id)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
