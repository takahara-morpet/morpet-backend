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
		// エラーの種類に応じて適切なレスポンスを返す
		if err.Error() == "post not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
		}
		// その他のエラーの場合は、500 Internal Server Error を返す
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// 正常な場合、ポスト情報を JSON で返す
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

	// 成功時にはステータスコード 200 を返す
	return c.NoContent(http.StatusOK)
}
