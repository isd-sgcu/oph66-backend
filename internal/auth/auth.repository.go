package auth

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(id uint, user *User) (*User, error)
	GetUserByEmail(user *User, email string) (*User, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) CreateUser(user *User) (*User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repositoryImpl) UpdateUser(id uint, user *User) (*User, error) {
	user.ID = id
	if err := r.db.Model(&user).Association("DesiredRounds").Replace(&user.DesiredRounds); err != nil {
		return nil, err
	}

	if err := r.db.Model(&user).Association("InterestedFaculties").Replace(&user.InterestedFaculties); err != nil {
		return nil, err
	}

	if err := r.db.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repositoryImpl) GetUserByEmail(user *User, email string) (*User, error) {
	if err := r.db.Where("email = ?", email).Preload("DesiredRounds").Preload("InterestedFaculties").First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
