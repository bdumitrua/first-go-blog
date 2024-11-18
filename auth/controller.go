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

func (uc *Controller) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	default:
		http.Error(w, "Route not found", http.StatusNotFound)
	}
}

func (uc *Controller) Login(w http.ResponseWriter, r *http.Request) {
	req := MakeLoginRequest(w, r)

	loginDto, err := req.ToDto()
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
	}

	token, err := uc.service.Login(loginDto)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadGateway)
	}

	json.NewEncoder(req.Writer()).Encode(token)
}

func (uc *Controller) Register(w http.ResponseWriter, r *http.Request) {
	req := MakeRegisterRequest(w, r)

	userCreateDto, err := req.ToDto()
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
	}

	message, err := uc.service.Register(userCreateDto)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadGateway)
	}

	json.NewEncoder(req.Writer()).Encode(message)
}

func (uc *Controller) Refresh(w http.ResponseWriter, r *http.Request) {
	req := utils.MakeRequest(w, r)

	token, err := GetJwtToken(req)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadRequest)
		return
	}

	newToken, err := uc.service.Refresh(token)
	if err != nil {
		http.Error(req.Writer(), err.Error(), http.StatusBadGateway)
	}

	json.NewEncoder(req.Writer()).Encode(newToken)
}
