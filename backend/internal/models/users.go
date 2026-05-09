package models

import "time"

type User struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"-" db:"password_hash"`

	DisplayName *string `json:"display_name" db:"display_name"`
	Bio         *string `json:"bio" db:"bio"`
	AvatarURL   *string `json:"avatar_url" db:"avatar_url"`
	IsPrivate   bool    `json:"is_private" db:"is_private"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
