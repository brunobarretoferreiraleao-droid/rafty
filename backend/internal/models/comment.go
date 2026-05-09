package models

import "time"

type Comment struct {
	ID              string     `db:"id" json:"id"`
	PostID          string     `db:"post_id" json:"post_id"`
	AuthorID        string     `db:"author_id" json:"author_id"`
	ParentCommentID *string    `db:"parent_comment_id" json:"parent_comment_id,omitempty"`
	Content         string     `db:"content" json:"content"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
