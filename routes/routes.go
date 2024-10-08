package routes

import (
	"morpet-backend/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.GET("/users/:id", handlers.GetUserById)

	e.POST("/replys", handlers.CreateReply)
	e.GET("/replys/:id", handlers.GetReplyById)
	e.GET("/replys/:post_id", handlers.GetRepliesByPostID)
	e.GET("/replys", handlers.GetReplys)

	e.GET("/posts", handlers.GetPosts)
	e.GET("/posts/:id", handlers.GetPostById)
	e.POST("/posts", handlers.CreatePost)
	e.PUT("/posts/:id/percentage", handlers.UpdatePostPercentage)
}
