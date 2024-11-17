package main

import (
	"first-blog-api/posts"
	"fmt"
	"net/http"
)

func main() {
	postRepo := posts.NewRepository()
	postService := posts.NewService(postRepo)
	postController := posts.NewController(postService)

	http.HandleFunc("/posts", postController.HandleRoutes)
	http.HandleFunc("/posts/", postController.HandleRoutes)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
