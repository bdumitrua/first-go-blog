package posts

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (pc *Controller) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) >= 3 && parts[2] != "" {
		postId, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Post ID must be integer", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			pc.GetPostById(w, postId)
		case http.MethodPut:
			pc.UpdatePost(w, r, postId)
		case http.MethodDelete:
			pc.DeletePost(w, postId)
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			pc.GetAll(w)
		case http.MethodPost:
			pc.CreatePost(w, r)
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	}
}

func (pc *Controller) GetAll(w http.ResponseWriter) {
	posts, err := pc.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (pc *Controller) GetPostById(w http.ResponseWriter, postId int) {
	post, err := pc.service.GetById(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (pc *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		http.Error(w, "Invalid post data", http.StatusBadRequest)
		return
	}

	mes, err := pc.service.CreatePost(newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(mes)
}

func (pc *Controller) UpdatePost(w http.ResponseWriter, r *http.Request, postId int) {
	var updatedPost Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Invalid update post data", http.StatusBadRequest)
		return
	}

	mes, err := pc.service.UpdatePost(updatedPost, postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(mes)
}

func (pc *Controller) DeletePost(w http.ResponseWriter, postId int) {
	mes, err := pc.service.DeletePost(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(mes)
}
