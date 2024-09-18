package dto

type CreateReplyRequest struct {
	PostID   int    `json:"postId" validate:"required"`   // リプライが関連するポストのID
	SenderID int    `json:"senderId" validate:"required"` // リプライを送信するユーザーのID
	Content  string `json:"content" validate:"required"`  // リプライの内容
}
