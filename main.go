package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts = []Post{
	{ID: 1, Title: "First title", Content: "First content"},
	{ID: 2, Title: "Second title", Content: "Second content"},
}

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling '/posts' url")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
