package posts

import (
	"errors"
	"fmt"
)

type Repository interface {
	GetAll() (*[]Post, error)
	GetById(postId int) (*Post, error)
	CreatePost(newPost Post) (string, error)
	UpdatePost(newPost Post, postId int) (string, error)
	DeletePost(postId int) (string, error)
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
	newPost.ID = r.posts[len(r.posts)-1].ID + 1
	r.posts = append(r.posts, newPost)

	return "Post created successfully", nil
}

func (r *repositoryImpl) UpdatePost(updatedPost Post, postId int) (string, error) {
	for i, post := range r.posts {
		if post.ID == postId {
			r.posts[i].Title = updatedPost.Title
			r.posts[i].Content = updatedPost.Content

			return "Post updated successfully", nil
		}
	}

	return "Post with id " + fmt.Sprint(postId) + " not found", errors.New("Post not found")
}

func (r *repositoryImpl) DeletePost(postId int) (string, error) {
	for i, post := range r.posts {
		if post.ID == postId {
			if len(r.posts) == i+1 {
				r.posts = posts[:i]
			} else {
				r.posts = append(posts[:i], posts[i+1:]...)
			}

			return "Post delete successfully", nil
		}
	}

	return "Post with id " + fmt.Sprint(postId) + " not found", errors.New("Post not found")
}
