package main

import (
	"first-blog-api/auth"
	"first-blog-api/db"
	"first-blog-api/middleware"
	"first-blog-api/posts"
	"first-blog-api/users"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	// Роуты для posts
	http.Handle("/posts/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) >= 3 && parts[2] != "" {
			_, err := strconv.Atoi(parts[2])
			if err != nil {
				http.Error(w, "Post ID must be integer", http.StatusBadRequest)
				return
			}

			switch r.Method {
			case http.MethodGet:
				postController.HandleRoutes(w, r)
			case http.MethodPut:
				middleware.AuthMiddleware(http.HandlerFunc(postController.HandleRoutes)).ServeHTTP(w, r)
			case http.MethodDelete:
				middleware.AuthMiddleware(http.HandlerFunc(postController.HandleRoutes)).ServeHTTP(w, r)
			default:
				http.Error(w, "Route not found", http.StatusNotFound)
			}
		} else {
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	}))

	http.Handle("/posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			postController.HandleRoutes(w, r)
		case http.MethodPost:
			middleware.AuthMiddleware(http.HandlerFunc(postController.HandleRoutes)).ServeHTTP(w, r)
		default:
			http.Error(w, "Route not found", http.StatusNotFound)
		}
	}))

	usersRoutes := http.HandlerFunc(userController.HandleRoutes)
	http.Handle("/users", middleware.AuthMiddleware(usersRoutes))

	http.HandleFunc("/login", authController.Login)
	http.HandleFunc("/register", authController.Register)
	http.Handle("/refresh", middleware.AuthMiddleware(http.HandlerFunc(authController.Refresh)))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
