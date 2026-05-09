package repositories

import (
	"context"
	"database/sql"
	"rafty/internal/models"
)

type PostRepository struct {
	DB *sql.DB
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (author_id, title, content, visibility)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRowContext(
		ctx,
		query,
		post.AuthorID,
		post.Title,
		post.Content,
		post.Visibility,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) GetFeedPosts(ctx context.Context, userID string, limit, offset int) ([]models.Post, error) {
	query := `
		SELECT id, author_id, title, content, visibility, created_at, updated_at
		FROM posts
		WHERE deleted_at IS NULL
		AND (
			author_id = $1

			OR visibility = 'public'

			OR (
				visibility = 'friends'
				AND EXISTS (
					SELECT 1 FROM friends
					WHERE user_id1 = LEAST(posts.author_id, $1)
					  AND user_id2 = GREATEST(posts.author_id, $1)
				)
			)

			OR (
				visibility = 'custom'
				AND EXISTS (
					SELECT 1 FROM post_visibility_rules pvr
					WHERE pvr.post_id = posts.id
					  AND pvr.user_id = $1
					  AND pvr.can_view = true
				)
			)
		)
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.DB.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.Visibility,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// DeletePost performs a soft delete by setting the DeletedAt timestamp for the specified post ID.
func (r *PostRepository) DeletePost(ctx context.Context, postID string, userID string) error {
	query := `
		UPDATE posts
		SET deleted_at = NOW()
		WHERE id = $1 AND author_id = $2
	`
	res, err := r.DB.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
