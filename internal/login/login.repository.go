package login

import (
	"gorm.io/gorm"
)

type Repository interface {
	GoogleLogin() (string, error)
	GoogleCallback() (string, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) GoogleLogin() (string, error) {
	return "", nil
}

func (r *repositoryImpl) GoogleCallback() (string, error) {
	return "", nil
}
