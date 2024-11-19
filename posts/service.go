package posts

type Service interface {
	GetAll() ([]Post, error)
	GetById(postId int) (*Post, error)
	CreatePost(newPostDto *PostDTO, userId int) (string, error)
	UpdatePost(updatePostDto *PostDTO, postId int, userId int) (string, error)
	DeletePost(postId int, userId int) (string, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) GetAll() ([]Post, error) {
	return s.repo.GetAll()
}

func (s *serviceImpl) GetById(postId int) (*Post, error) {
	return s.repo.GetById(postId)
}

func (s *serviceImpl) CreatePost(newPostDTO *PostDTO, userId int) (string, error) {
	return s.repo.CreatePost(newPostDTO.ToPost(), userId)
}

func (s *serviceImpl) UpdatePost(updatePostDto *PostDTO, postId int, userId int) (string, error) {
	return s.repo.UpdatePost(updatePostDto.ToPost(), postId, userId)
}

func (s *serviceImpl) DeletePost(postId int, userId int) (string, error) {
	return s.repo.DeletePost(postId, userId)
}
