package posts

type Service interface {
	GetAll() (*[]Post, error)
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
