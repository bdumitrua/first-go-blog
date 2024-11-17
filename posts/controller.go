package posts

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts, err = c.service.GetAll()
	if err != nil {
		http.Error(w, "Something gone wrong while requested posts", 500)
	}

	json.NewEncoder(w).Encode(posts)
}
