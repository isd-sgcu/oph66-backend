package staff

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCheckin(checkin *model.AttendeeCheckin) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) CreateCheckin(checkin *model.AttendeeCheckin) error {
	return r.db.Model(checkin).Create(checkin).Error
}
