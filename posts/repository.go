package posts

import (
	"errors"
	"fmt"
)

type Repository interface {
	GetAll() (*[]Post, error)
	GetById(postId int) (*Post, error)
	CreatePost(newPost Post) (string, error)
}

type repositoryImpl struct {
	posts []Post
}

func NewRepository() Repository {
	return &repositoryImpl{posts: posts}
}

func (r *repositoryImpl) GetAll() (*[]Post, error) {
	return &r.posts, nil
}

func (r *repositoryImpl) GetById(postId int) (*Post, error) {
	for _, post := range r.posts {
		if post.ID == postId {
			return &post, nil
		}
	}

	return nil, errors.New("Post with id " + fmt.Sprint(postId) + " not found")
}

func (r *repositoryImpl) CreatePost(newPost Post) (string, error) {
	newPost.ID = len(r.posts) + 1
	r.posts = append(r.posts, newPost)

	return "Post created successfully", nil
}
