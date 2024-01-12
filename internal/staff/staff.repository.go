package staff

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateFacultyCheckin(checkin *model.AttendeeFacultyCheckin) error
	CreateCentralCheckin(checkin *model.AttendeeCentralCheckin) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) CreateFacultyCheckin(checkin *model.AttendeeFacultyCheckin) error {
	return r.db.Model(checkin).Create(checkin).Error
}

func (r *repositoryImpl) CreateCentralCheckin(checkin *model.AttendeeCentralCheckin) error {
	return r.db.Model(checkin).Create(checkin).Error
}
