package main

import (
	"morpet-backend/routes"
	"morpet-backend/config"

	"github.com/labstack/echo/v4"
)

func main() {

	config.InitDB()
	defer config.CloseDB()

	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.GET("/users", handlers.GetUsers)
	routes.InitRoutes(e)
	
	e.Logger.Fatal(e.Start(":8080"))
}
