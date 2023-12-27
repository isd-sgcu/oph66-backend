package event

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllEvents(results *[]model.Event) error
	GetEventById(result *model.Event, eventId string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) GetAllEvents(results *[]model.Event) error {
	return r.db.Model(&model.Event{}).Omit("description_th, description_en").Preload("Schedules").Preload("Faculty").Find(results).Error
}

func (r *repositoryImpl) GetEventById(result *model.Event, eventId string) error {
	return r.db.Model(&model.Event{}).Preload("Schedules").Preload("Faculty").First(result, "id = ?", eventId).Error
}
