package models

import "time"

type Post struct {
	ID         string     `db:"id" json:"id"`
	AuthorID   string     `db:"author_id" json:"author_id"`
	Title      *string    `db:"title" json:"title,omitempty"`
	Content    *string    `db:"content" json:"content,omitempty"`
	Visibility string     `db:"visibility" json:"visibility"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type PostAttachment struct {
	ID             string    `db:"id" json:"id"`
	PostID         string    `db:"post_id" json:"post_id"`
	FileURL        string    `db:"file_url" json:"file_url"`
	FileName       *string   `db:"file_name" json:"file_name,omitempty"`
	MimeType       *string   `db:"mime_type" json:"mime_type,omitempty"`
	FileSize       *int64    `db:"file_size" json:"file_size,omitempty"`
	AttachmentType string    `db:"attachment_type" json:"attachment_type"`
	DisplayOrder   int       `db:"display_order" json:"display_order"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
