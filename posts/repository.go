package posts

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	GetAll() ([]Post, error)
	GetById(postId int) (*Post, error)
	CreatePost(newPost Post) (string, error)
	UpdatePost(newPost Post, postId int) (string, error)
	DeletePost(postId int) (string, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetAll() ([]Post, error) {
	posts := []Post{}

	rows, err := r.db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *repositoryImpl) GetById(postId int) (*Post, error) {
	var post Post

	err := r.db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", postId).Scan(&post.ID, &post.Title, &post.Content)
	if err == sql.ErrNoRows {
		return nil, errors.New("Post not found")
	} else if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *repositoryImpl) CreatePost(newPost Post) (string, error) {
	result, err := r.db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", newPost.Title, newPost.Content)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()
	return fmt.Sprintf("Post created successfully with ID %d", id), nil
}

func (r *repositoryImpl) UpdatePost(updatedPost Post, postId int) (string, error) {
	result, err := r.db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", updatedPost.Title, updatedPost.Content, postId)
	if err != nil {
		return "", err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return "", errors.New("Post not found")
	}

	return "Post updated successfully", nil
}

func (r *repositoryImpl) DeletePost(postId int) (string, error) {
	result, err := r.db.Exec("DELETE FROM posts WHERE id = ?", postId)
	if err != nil {
		return "", err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return "", errors.New("Post not found")
	}

	return "Post deleted successfully", nil
}
