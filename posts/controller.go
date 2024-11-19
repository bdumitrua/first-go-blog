package posts

import (
	"encoding/json"
	"first-blog-api/auth"
	"first-blog-api/utils"
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
			pc.GetPostById(utils.MakeRequest(w, r), postId)
		case http.MethodPut:
			pc.UpdatePost(MakeUpdateRequest(w, r), postId)
		case http.MethodDelete:
			pc.DeletePost(utils.MakeRequest(w, r), postId)
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			pc.GetAll(utils.MakeRequest(w, r))
		case http.MethodPost:
			pc.CreatePost(MakeCreateRequest(w, r))
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	}
}

func (pc *Controller) GetAll(req *utils.Request) {
	posts, err := pc.service.GetAll()
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(posts)
}

func (pc *Controller) GetPostById(req *utils.Request, postId int) {
	post, err := pc.service.GetById(postId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(req.Writer()).Encode(post)
}

func (pc *Controller) CreatePost(req *CreatePostRequest) {
	newPostDTO, err := req.ToDTO()
	if err != nil {
		return
	}

	userId, err := auth.GetUserId(&req.Request)
	if err != nil {
		return
	}

	mes, err := pc.service.CreatePost(newPostDTO, userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(mes)
}

func (pc *Controller) UpdatePost(req *UpdatePostRequest, postId int) {
	updatePostDTO, err := req.ToDTO()
	if err != nil {
		return
	}

	userId, err := auth.GetUserId(&req.Request)
	if err != nil {
		return
	}

	mes, err := pc.service.UpdatePost(updatePostDTO, postId, userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(mes)
}

func (pc *Controller) DeletePost(req *utils.Request, postId int) {
	userId, err := auth.GetUserId(req)
	if err != nil {
		return
	}

	mes, err := pc.service.DeletePost(postId, userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(mes)
}
