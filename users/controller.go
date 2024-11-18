package users

import (
	"encoding/json"
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

func (uc *Controller) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) >= 3 && parts[2] != "" {
		userId, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Post ID must be integer", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPut:
			uc.UpdateUser(MakeUpdateRequest(w, r), userId)
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			uc.GetMe(utils.MakeRequest(w, r))
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	}
}

func (uc *Controller) GetMe(req *utils.Request) {
	userId := 1

	user, err := uc.service.GetById(userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(user)
}

func (uc *Controller) GetUserById(req *utils.Request, userId int) {
	user, err := uc.service.GetById(userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(req.Writer()).Encode(user)
}

func (uc *Controller) UpdateUser(req *UpdateUserRequest, userId int) {
	updateUserDTO, err := req.ToDTO()
	if err != nil {
		return
	}

	mes, err := uc.service.UpdateUser(updateUserDTO, userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(mes)
}
