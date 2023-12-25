package auth

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(id uint, user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
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
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repositoryImpl) UpdateUser(id uint, user *User) (*User, error) {
    tx := r.db.Begin()

    var existingUser User
    if err := tx.Where("id = ?", id).Preload("InterestedFaculties").Preload("DesiredRounds").First(&existingUser).Error; err != nil {
        tx.Rollback()
        return nil, err
    }

	if err := tx.Where("user_id = ?", id).Delete(&InterestedFaculties{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Where("user_id = ?", id).Delete(&DesiredRounds{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

    existingUser.FirstName = user.FirstName
    existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.School = user.School
	existingUser.BirthDate = user.BirthDate
	existingUser.Address = user.Address
	existingUser.FromAbroad = user.FromAbroad
	existingUser.Allergy = user.Allergy
	existingUser.MedicalCondition = user.MedicalCondition
	existingUser.JoinCUReason = user.JoinCUReason
	existingUser.NewsSource = user.NewsSource
	existingUser.Status = user.Status
	existingUser.Grade = user.Grade

    existingUser.InterestedFaculties = user.InterestedFaculties
    existingUser.DesiredRounds = user.DesiredRounds

    if err := tx.Save(&existingUser).Error; err != nil {
        tx.Rollback()
        return nil, err
    }

    if err := tx.Commit().Error; err != nil {
        return nil, err
    }

    return &existingUser, nil
}

func (r *repositoryImpl) GetUserByEmail(email string) (*User, error) {
	var user *User
    if err := r.db.Where("email = ?", email).Preload("InterestedFaculties").Preload("DesiredRounds").First(&user).Error; err != nil {
        return nil, err
    }

    return user, nil
}
