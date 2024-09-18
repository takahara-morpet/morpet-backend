package services

import (
	"database/sql"
	"fmt"
	"morpet-backend/config"
	"morpet-backend/handlers/dto"
	"morpet-backend/models"
)

func GetReplyById(id string) (models.Reply, error) {
	query := `
	SELECT id, post_id, sender_id, content, created_at, updated_at
	FROM replys
	WHERE id = $1
`
	row := config.DB.QueryRow(query, id)

	var reply models.Reply
	if err := row.Scan(
		&reply.ID, &reply.PostID, &reply.SenderID, &reply.Content, &reply.CreatedAt, &reply.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return models.Reply{}, fmt.Errorf("reply not found")
		}
		return models.Reply{}, fmt.Errorf("failed to scan reply: %w", err)
	}
	return reply, nil
}

func GetAllReplys() ([]models.Reply, error) {
	query := `
	SELECT id, post_id, sender_id, content, created_at, updated_at
	FROM replys
`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve replys: %w", err)
	}

	defer rows.Close()

	var replys []models.Reply
	for rows.Next() {
		var reply models.Reply
		if err := rows.Scan(&reply.ID, &reply.PostID, &reply.SenderID, &reply.Content, &reply.CreatedAt, &reply.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan reply: %w", err)
		}

		replys = append(replys, reply)
	}
	return replys, nil

}

func CreateReply(reply *dto.CreateReplyRequest) (int, error) {
	query := `
		INSERT INTO replys (post_id, sender_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`

	// 新しいリプライのIDを格納する変数
	var replyID int

	// クエリを実行して新しいリプライを挿入し、挿入されたIDを取得
	err := config.DB.QueryRow(query, reply.PostID, reply.SenderID, reply.Content).Scan(&replyID)
	if err != nil {
		return 0, fmt.Errorf("failed to create reply: %w", err)
	}

	return replyID, nil
}
