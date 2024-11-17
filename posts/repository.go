package posts

type Repository interface {
	GetAll() (*[]Post, error)
}

type repositoryImpl struct {
	posts []Post
}

func NewRepository() Repository {
	return &repositoryImpl{posts: posts}
}

func (r *repositoryImpl) GetAll() (*[]Post, error) {
	return &r.posts, nil
}
