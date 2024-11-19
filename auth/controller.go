package auth

import (
	"encoding/json"
	"first-blog-api/utils"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ac *Controller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := MakeLoginRequest(w, r)

	loginDto, err := req.ToDto()
	if err != nil {
		return
	}

	token, err := ac.service.Login(loginDto)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(req.Writer()).Encode(token)
}

func (ac *Controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := MakeRegisterRequest(w, r)

	userCreateDto, err := req.ToDto()
	if err != nil {
		return
	}

	message, err := ac.service.Register(userCreateDto)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadGateway)
		return
	}

	json.NewEncoder(req.Writer()).Encode(message)
}

func (ac *Controller) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := utils.MakeRequest(w, r)

	token, err := GetJwtToken(req)
	if err != nil {
		return
	}

	newToken, err := ac.service.Refresh(token)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadGateway)
		return
	}

	json.NewEncoder(req.Writer()).Encode(newToken)
}
