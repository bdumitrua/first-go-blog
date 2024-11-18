package main

import (
	"first-blog-api/auth"
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

	authRepo := auth.NewRepository(db.DB)
	authService := auth.NewService(authRepo)
	authController := auth.NewController(authService)

	http.HandleFunc("/posts", postController.HandleRoutes)
	http.HandleFunc("/posts/", postController.HandleRoutes)

	http.HandleFunc("/users", userController.HandleRoutes)

	http.HandleFunc("/login", authController.Login)
	http.HandleFunc("/register", authController.Register)
	http.HandleFunc("/refresh", authController.Refresh)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
