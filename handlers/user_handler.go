package handlers

import (
	"morpet-backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Owataaaaa",
		})
	}
	return c.JSON(http.StatusOK, users)
}
