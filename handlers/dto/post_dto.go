package dto

type CreatePostRequest struct {
	UserID   int    `json:"userId" validate:"required"`   // ポストを作成するユーザーのID
	Content  string `json:"content" validate:"required"`  // ポストの内容
	Category string `json:"category" validate:"required"` // ポストのカテゴリ
}
