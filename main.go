package main

import (
	"morpet-backend/config"
	"morpet-backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
	"os"
	"net/http"
)

func main() {

	config.InitDB()
	defer config.CloseDB()

	e := echo.New()
	// swaggerだけ許可されるようにしている
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	routes.InitRoutes(e)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not set
    }

    // Start the server
    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
