package handlers

import (
	"log"
	"morpet-backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Owataaaaa",
		})
	}
	return c.JSON(http.StatusOK, users)
}
