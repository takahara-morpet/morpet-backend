package routes

import (
	"morpet-backend/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/users", handlers.GetUsers)
	e.POST("/user", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
}
