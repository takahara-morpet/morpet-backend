package dto

// dto.CreateUserRequest
type CreateUserRequest struct {
	Name            string `json:"name"`
	Age             int    `json:"age"`
	Profile         string `json:"profile"`
	ProfileImageUrl string `json:"profileImageUrl"`
	Gender          string `json:"gender"`
	Mbti            string `json:"mbti"`
}

// dto.UpdateUserRequest
type UpdateUserRequest struct {
	Age             int    `json:"age"`
	Profile         string `json:"profile"`
	ProfileImageUrl string `json:"profileImageUrl"`
	Gender          string `json:"gender"`
	Mbti            string `json:"mbti"`
}
