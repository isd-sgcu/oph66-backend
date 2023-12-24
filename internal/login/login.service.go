package login

type Service interface {
	GoogleLogin() (string, error)
	GoogleCallback() (string, error)
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo}
}

type serviceImpl struct {
	repo Repository
}

func (s *serviceImpl) GoogleLogin() (string, error) {
	return s.repo.GoogleLogin()
}

func (s *serviceImpl) GoogleCallback() (string, error) {
	return s.repo.GoogleCallback()
}
