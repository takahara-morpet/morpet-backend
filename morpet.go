package main

import (
	"morpet-backend/config"
	"morpet-backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitDB()
	defer config.CloseDB()

	e := echo.New()
	// swaggerだけ許可されるようにしている
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.GET("/users", handlers.GetUsers)
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
