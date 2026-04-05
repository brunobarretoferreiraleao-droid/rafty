package repositories

import (
	"database/sql"
	"rafty/internal/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

// CreatePost inserts a new post into the database and returns the created post with its ID.
func (r *PostRepository) CreatePost(post *models.Post) error {
	query := `
		INSERT INTO posts (author_id, title, content, visibility)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(
		query,
		post.AuthorID,
		post.Title,
		post.Content,
		post.Visibility,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}

// GetPostByID retrieves a post by its ID, ensuring it is not soft-deleted.
func (r *PostRepository) GetPostByID(postID string) (*models.Post, error) {
	query := `
		SELECT id, author_id, title, content, visibility, created_at, updated_at, deleted_at
		FROM posts
		WHERE id = $1 AND deleted_at IS NULL
	`

	var post models.Post
	err := r.DB.QueryRow(query, postID).Scan(
		&post.ID,
		&post.AuthorID,
		&post.Title,
		&post.Content,
		&post.Visibility,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetFeedPosts retrieves a list of posts for the feed, applying pagination and visibility filters.
func (r *PostRepository) GetFeedPosts(userID string, limit, offset int) ([]models.Post, error) {
	query := `
		SELECT id, author_id, title, content, visibility, created_at, updated_at, deleted_at
		FROM posts
		WHERE deleted_at IS NULL
		  AND (visibility = 'public' OR author_id = $1)
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.DB.Query(query, userID, limit, offset)
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
			&post.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// DeletePost performs a soft delete by setting the DeletedAt timestamp for the specified post ID.
func (r *PostRepository) DeletePost(postID string) error {
	query := `
		UPDATE posts
		SET deleted_at = NOW()
		WHERE id = $1
	`
	_, err := r.DB.Exec(query, postID)
	return err
}
