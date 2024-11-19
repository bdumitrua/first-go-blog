package users

import (
	"encoding/json"
	"first-blog-api/auth"
	"first-blog-api/utils"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (uc *Controller) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		uc.GetMe(utils.MakeRequest(w, r))
	case http.MethodPut:
		uc.UpdateUser(MakeUpdateRequest(w, r))
	default:
		http.Error(w, "Route not found", http.StatusNotFound)
	}

}

func (uc *Controller) GetUserById(req *utils.Request, userId int) {
	user, err := uc.service.GetById(userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(req.Writer()).Encode(user)
}

func (uc *Controller) GetMe(req *utils.Request) {
	userId, err := auth.GetUserId(req)
	if err != nil {
		return
	}

	user, err := uc.service.GetById(userId)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(user)
}

func (uc *Controller) UpdateUser(req *UpdateUserRequest) {
	userId, err := auth.GetUserId(&req.Request)
	if err != nil {
		return
	}

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
