package event

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllEvents(results *[]Event) error
	GetEventById(result *Event, eventId string) error
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
	return r.db.Model(&Event{}).Omit("description_th, description_en").Preload("Schedules").Preload("Faculty").Find(results).Error
}

func (r *repositoryImpl) GetEventById(result *Event, eventId string) error {
	return r.db.Model(&Event{}).Preload("Schedules").Preload("Faculty").First(result, "id = ?", eventId).Error
}
