package auth

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(user *User, email string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) CreateUser(user *User) error {
	return r.db.Create(&user).Error
}

func (r *repositoryImpl) GetUserByEmail(user *User, email string) error {
	return r.db.Preload("DesiredRounds").Preload("DesiredRounds.Round").Preload("InterestedFaculties").Preload("InterestedFaculties.Faculty").Preload("InterestedFaculties.Department").Preload("InterestedFaculties.Section").Where("email = ?", email).First(&user).Error
}