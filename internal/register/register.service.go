package register

import ()

type Service interface {
	CreateUser(user *User) error
	GetUserById(id uint) (*User, error)
	GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, error)
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo}
}

type serviceImpl struct {
	repo Repository
}

func (s *serviceImpl) CreateUser(user *User) error {
	print(user)
	return s.repo.Create(user)
}

func (s *serviceImpl) GetUserById(id uint) (*User, error) {
	return s.repo.GetById(id)
}

func (s *serviceImpl) GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, error) {
	return s.repo.GetInterestedFacultiesByUserId(id)
}
