package services

import (
	"morpet-backend/models"
	"morpet-backend/config"
	"log"
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
