package handlers

import (
	"fmt"
	"log"
	"morpet-backend/handlers/dto"
	"morpet-backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetReplyById(c echo.Context) error {
	id := c.Param("id")

	// サービス層からユーザーを取得
	reply, err := services.GetReplyById(id)
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if err.Error() == "reply not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Reply not found"})
		}
		// その他のエラーの場合は、500 Internal Server Error を返す
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// 正常な場合、ユーザー情報を JSON で返す
	return c.JSON(http.StatusOK, reply)
}
func GetRepliesByPostID(c echo.Context) error {
	// パスパラメータから post_id を取得
	postID := c.Param("post_id")

	// サービス層から post_id に関連する返信を取得
	replies, err := services.GetRepliesByPostID(postID)
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if err.Error() == "replies not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Replies not found"})
		}
		// その他のエラーの場合は、500 Internal Server Error を返す
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// 正常な場合、返信情報を JSON で返す
	return c.JSON(http.StatusOK, replies)
}

func GetReplys(c echo.Context) error {
	replys, err := services.GetAllReplys()
	if err != nil {
		log.Printf("Failed to get replys: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Owataaaaa",
		})
	}
	return c.JSON(http.StatusOK, replys)
}

func CreateReply(c echo.Context) error {
	req := &dto.CreateReplyRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	replyID, err := services.CreateReply(req)
	if err != nil {
		return fmt.Errorf("replyの作成に失敗しました:%s", err)
	}
	return c.JSON(http.StatusOK, replyID)

}
