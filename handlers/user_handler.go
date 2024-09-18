package handlers

import (
	"fmt"
	"log"
	"morpet-backend/handlers/dto"
	"morpet-backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserById(c echo.Context) error {
	id := c.Param("id")

	// サービス層からユーザーを取得
	user, err := services.GetUserById(id)
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		// その他のエラーの場合は、500 Internal Server Error を返す
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// 正常な場合、ユーザー情報を JSON で返す
	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Owataaaaa",
		})
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	req := &dto.CreateUserRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := services.CreateUser(req); err != nil {
		return fmt.Errorf("ユーザの作成に失敗しました:%s", err)
	}
	return c.NoContent(http.StatusOK)

}

func UpdateUser(c echo.Context) error {
	req := &dto.UpdateUserRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	id := c.Param("id")
	err := services.UpdateUser(id, req)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}
	return c.NoContent(http.StatusOK)
}
