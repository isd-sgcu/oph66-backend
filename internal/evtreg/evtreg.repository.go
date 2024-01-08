package evtreg

import (
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetUserWithEventRegistrationByEmail(user *model.User, email string) error
	GetScheduleById(schedule *model.Schedule, scheduleId int) error
	RegisterEvent(evtreg *model.EventRegistration) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

type repositoryImpl struct {
	db *gorm.DB
}

func (r *repositoryImpl) GetUserWithEventRegistrationByEmail(user *model.User, email string) error {
	return r.db.Model(user).Preload("RegisteredEvents").Preload("RegisteredEvents.Schedule").Where("email = ?", email).First(&user).Error
}

func (r *repositoryImpl) GetScheduleById(schedule *model.Schedule, scheduleId int) error {
	return r.db.Model(schedule).Where("id = ?", scheduleId).First(schedule).Error
}

func (r *repositoryImpl) RegisterEvent(evtreg *model.EventRegistration) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var schedule model.Schedule

		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Model(&schedule).Preload("Event").Where("id = ?", evtreg.ScheduleId).Find(&schedule).Error; err != nil {
			return err
		}

		if schedule.CurrentAttendee >= schedule.Event.MaxCapacity {
			return apperror.ScheduleFull
		}

		schedule.CurrentAttendee++
		tx.Save(&schedule)

		if err := tx.Create(&evtreg).Error; err != nil {
			return err
		}

		return nil
	})
}
