package posts

type Service interface {
	GetAll() (*[]Post, error)
	GetById(postId int) (*Post, error)
	CreatePost(newPost Post) (string, error)
	UpdatePost(UpdatePost Post, postId int) (string, error)
	DeletePost(postId int) (string, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) GetAll() (*[]Post, error) {
	return s.repo.GetAll()
}

func (s *serviceImpl) GetById(postId int) (*Post, error) {
	return s.repo.GetById(postId)
}

func (s *serviceImpl) CreatePost(newPost Post) (string, error) {
	return s.repo.CreatePost(newPost)
}

func (s *serviceImpl) UpdatePost(newPost Post, postId int) (string, error) {
	return s.repo.UpdatePost(newPost, postId)
}

func (s *serviceImpl) DeletePost(postId int) (string, error) {
	return s.repo.DeletePost(postId)
}
