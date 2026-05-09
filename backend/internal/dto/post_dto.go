package dto

type CreatePostRequest struct {
	Title      *string `json:"title"`
	Content    *string `json:"content"`
	Visibility string  `json:"visibility"`
}

type UpdatePostRequest struct {
	Title      *string `json:"title"`
	Content    *string `json:"content"`
	Visibility *string `json:"visibility"`
}
