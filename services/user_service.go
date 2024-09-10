package services

import "morpet-backend/models"

func GetAllUsers() ([]models.User, error) {
	users := []models.User{
		{Id: 1, Name: "name", Age: 5},
		{Id: 21, Name: "sfvdsname", Age: 5},
	}

	return users, nil
}

// func GetAllUsers() ([]models.User, error) {
// 	rows, err := config.DB.Query("SELECT id, name, age FROM users")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var users []models.User
// 	for rows.Next() {
// 		var user models.User
// 		if err := rows.Scan(
// 			&user.Id,
// 			&user.Name,
// 			&user.Age,
// 		); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }
