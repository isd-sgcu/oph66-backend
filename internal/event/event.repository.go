package event

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllEvents(result *[]Event) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}
func (r *repositoryImpl) GetAllEvents(results *[]Event) error {
	return r.db.Model(&Event{}).Omit("description").Preload("Faculty").Find(results).Error
}
