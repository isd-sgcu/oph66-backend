package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(user *model.User, email string) error
	GetUserById(model *model.User, id int) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) CreateUser(user *model.User) error {
	return r.db.Model(user).Create(&user).Error
}

func (r *repositoryImpl) GetUserByEmail(user *model.User, email string) error {
	return r.db.
		Model(user).
		Preload("RegisteredEvents").
		Preload("NewsSourceUsers").
		Preload("InterestedFaculties").
		Preload("InterestedFaculties.Faculty").
		Preload("InterestedFaculties.Department").
		Preload("InterestedFaculties.Section").
		Preload("VisitingFaculties").
		Preload("VisitingFaculties.Faculty").
		Preload("VisitingFaculties.Department").
		Preload("VisitingFaculties.Section").
		Preload("RegisteredEvents.Schedule").
		Where("email = ?", email).
		First(&user).Error
}

func (r *repositoryImpl) GetUserById(user *model.User, id int) error {
	return r.db.Model(user).First(&user, "id = ?", id).Error
}
