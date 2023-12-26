package auth

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	UpdateUser(user *User) error
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

func (r *repositoryImpl) UpdateUser(user *User) error {
	if err := r.db.Model(&user).Association("DesiredRounds").Replace(&user.DesiredRounds); err != nil {
		return err
	}

	if err := r.db.Model(&user).Association("InterestedFaculties").Replace(&user.InterestedFaculties); err != nil {
		return err
	}

	return r.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
}

func (r *repositoryImpl) GetUserByEmail(user *User, email string) error {
	return r.db.Preload("DesiredRounds").Preload("DesiredRounds.Round").Preload("InterestedFaculties").Preload("InterestedFaculties.Faculty").Preload("InterestedFaculties.Department").Preload("InterestedFaculties.Section").Where("email = ?", email).First(&user).Error
}
