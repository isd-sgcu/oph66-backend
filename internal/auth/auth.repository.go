package auth

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(result *User) error
	GetByEmail(email string) (*User, error)
	GetInterestedFacultiesByUserId(id uint) ([]InterestedFaculties, error)
	GetDesiredRoundsByUserId(id uint) ([]DesiredRounds, error)
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

func (r *repositoryImpl) GetByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
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

func (r *repositoryImpl) GetDesiredRoundsByUserId(id uint) ([]DesiredRounds, error) {
	var desiredRounds []DesiredRounds
	if err := r.db.Where("user_id = ?", id).Find(&desiredRounds).Error; err != nil {
		return nil, err
	}
	return desiredRounds, nil
}