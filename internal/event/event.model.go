package event

import (
	"github.com/isd-sgcu/oph66-backend/internal/faculty"
	"github.com/isd-sgcu/oph66-backend/internal/schedule"
)

type Event struct {
	Id                  string              `gorm:"primaryKey" json:"id"`
	Name                string              `json:"name"`
	FacultyCode         int                 `gorm:"references:Code" json:"-"`
	Faculty             faculty.Faculty     `json:"faculty"`
	Department          string              `json:"department"`
	RequireRegistration bool                `json:"require_registration"`
	MaxCapacity         int                 `json:"max_capacity"`
	Schedules           []schedule.Schedule `json:"schedules" gorm:"foreignKey:event_id"`
	LocationEn          string              `json:"location_en"`
	LocationTh          string              `json:"location_th"`
	Description         string              `json:"description,omitempty"`
}

func (m Event) TableName() string {
	return "events"
}
