package main

import (
	"morpet-backend/config"
	"morpet-backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
	"os"
)

func main() {

	config.InitDB()
	defer config.CloseDB()

	e := echo.New()
	// swaggerだけ許可されるようにしている
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{
			"https://morpet-frontend.vercel.app",
			"http://localhost:8000",
			"http://localhost:3000",
		},
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	

	routes.InitRoutes(e)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not set
    }

    // Start the server
    log.Printf("Starting server on port %s", port)
    if err := e.Start(":"+port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
