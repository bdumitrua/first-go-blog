package posts

type Service interface {
	GetAll() (*[]Post, error)
	GetById(postId int) (*Post, error)
	CreatePost(newPostDto *PostDTO) (string, error)
	UpdatePost(updatePostDto *PostDTO, postId int) (string, error)
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

func (s *serviceImpl) CreatePost(newPostDTO *PostDTO) (string, error) {
	return s.repo.CreatePost(newPostDTO.ToPost())
}

func (s *serviceImpl) UpdatePost(updatePostDto *PostDTO, postId int) (string, error) {
	return s.repo.UpdatePost(updatePostDto.ToPost(), postId)
}

func (s *serviceImpl) DeletePost(postId int) (string, error) {
	return s.repo.DeletePost(postId)
}
