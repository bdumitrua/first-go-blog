package main

import (
	"first-blog-api/db"
	"first-blog-api/posts"
	"first-blog-api/users"
	"fmt"
	"net/http"
)

func main() {
	// Подключаемся к базе данных
	db.Connect()

	postRepo := posts.NewRepository(db.DB)
	postService := posts.NewService(postRepo)
	postController := posts.NewController(postService)

	userRepo := users.NewRepository(db.DB)
	userService := users.NewService(userRepo)
	userController := users.NewController(userService)

	http.HandleFunc("/posts", postController.HandleRoutes)
	http.HandleFunc("/posts/", postController.HandleRoutes)

	http.HandleFunc("/users", userController.HandleRoutes)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
