package users

type Service interface {
	GetById(userId int) (*User, error)
	UpdateUser(userUpdateDto *UserUpdateDTO, userId int) (string, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) GetById(userId int) (*User, error) {
	return s.repo.GetById(userId)
}

func (s *serviceImpl) UpdateUser(userUpdateDto *UserUpdateDTO, userId int) (string, error) {
	return s.repo.UpdateUser(userUpdateDto, userId)
}
