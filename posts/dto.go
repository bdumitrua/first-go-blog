package posts

type PostDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (dto *PostDTO) ToPost() Post {
	return Post{Title: dto.Title, Content: dto.Content}
}
