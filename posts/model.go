package posts

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts = []Post{
	{ID: 1, Title: "First title", Content: "First content"},
	{ID: 2, Title: "Second title", Content: "Second content"},
}
