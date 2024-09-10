package main

import (
	"morpet-backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.InitDB()
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.GET("/users", handlers.GetUsers)
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
