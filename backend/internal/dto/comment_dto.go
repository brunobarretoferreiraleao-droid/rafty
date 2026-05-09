package dto

type CreateCommentRequest struct {
	Content         string  `json:"content"`
	ParentCommentID *string `json:"parent_comment_id"`
}
