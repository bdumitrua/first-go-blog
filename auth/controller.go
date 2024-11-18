package auth

import (
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

func (uc *Controller) Login(req *LoginRequest) {
}

func (uc *Controller) Register(req *RegisterRequest) {
}

func (uc *Controller) Logout(req *utils.Request, token string) {
}

func (uc *Controller) Refresh(req *utils.Request, token string) {
}
