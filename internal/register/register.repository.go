package register

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetById(id uint) (*User, error)
	GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}
func (r *repositoryImpl) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repositoryImpl) GetById(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, error) {
	var interestedFaculties []InterestedFaculties
	if err := r.db.Where("user_id = ?", id).Find(&interestedFaculties).Error; err != nil {
		return nil, err
	}
	return interestedFaculties, nil
}
