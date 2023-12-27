package evtreg

import (
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetUserWithEventRegistrationByEmail(user *model.User, email string) error
	GetScheduleByID(schedule *model.Schedule, scheduleId int) error
	RegisterEvent(userID int, scheduleID int) error
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
	return r.db.Preload("RegisteredEvents").Preload("RegisteredEvents.Schedule").Where("email = ?", email).First(&user).Error
}

func (r *repositoryImpl) GetScheduleByID(schedule *model.Schedule, scheduleID int) error {
	return r.db.Where("id = ?", scheduleID).First(schedule).Error
}

func (r *repositoryImpl) RegisterEvent(userID int, scheduleID int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var schedule model.Schedule

		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Model(&schedule).Preload("Event").Where("id = ?", scheduleID).Find(&schedule).Error; err != nil {
			return err
		}

		if schedule.CurrentAttendee >= schedule.Event.MaxCapacity {
			return apperror.ScheduleFull
		}

		schedule.CurrentAttendee++
		tx.Save(&schedule)

		var reg model.EventRegistration
		reg.ScheduleId = scheduleID
		reg.UserID = userID

		if err := tx.Create(&reg).Error; err != nil {
			return err
		}

		return nil
	})
}
