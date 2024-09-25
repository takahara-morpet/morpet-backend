package handlers

import (
	"fmt"
	"log"
	"morpet-backend/handlers/dto"
	"morpet-backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPostById(c echo.Context) error {
	id := c.Param("id")

	// サービス層からポストを取得
	post, err := services.GetPostById(id)
	if err != nil {
		if err.Error() == "post not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {
	posts, err := services.GetAllPosts()
	if err != nil {
		log.Printf("Failed to get posts: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve posts",
		})
	}
	return c.JSON(http.StatusOK, posts)
}

func CreatePost(c echo.Context) error {
	req := &dto.CreatePostRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	// サービス層でポストを作成
	if _, err := services.CreatePost(req); err != nil {
		return fmt.Errorf("postの作成に失敗しました: %s", err)
	}

	return c.NoContent(http.StatusOK)
}

func UpdatePostPercentage(c echo.Context) error {
	type PercentageUpdateRequest struct {
		MalePercentage   int `json:"malePercentage"`
		FemalePercentage int `json:"femalePercentage"`
	}

	postID := c.Param("id")

	var req = &PercentageUpdateRequest{}
	if err := c.Bind(req); err != nil {
		log.Printf("Request body binding failed: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// パーセンテージのデータが正しく受け取れているかログに出力
	log.Printf("Received data - Male: %d, Female: %d\n", req.MalePercentage, req.FemalePercentage)

	err := services.UpdatePostPercentage(postID, req.MalePercentage, req.FemalePercentage)
	if err != nil {
		log.Printf("Failed to update percentages: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update percentages",
		})
	}

	return c.NoContent(http.StatusOK)
}
