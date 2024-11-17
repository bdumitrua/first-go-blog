package posts

type Service interface {
	GetAll() (*[]Post, error)
	GetById(postId int) (*Post, error)
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
