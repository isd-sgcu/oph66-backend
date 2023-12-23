package event

import (
	"github.com/isd-sgcu/oph66-backend/internal/bilingual_field"
	"github.com/isd-sgcu/oph66-backend/internal/faculty"
	"github.com/isd-sgcu/oph66-backend/internal/schedule"
)

type Event struct {
	Id                  string                     `json:"id" gorm:"primaryKey"`
	Name                bilingual_field.Bilingual  `json:"name" gorm:"embedded;embeddedPrefix:name_"`
	FacultyCode         int                        `json:"-" gorm:"references:Code"`
	Faculty             faculty.Faculty            `json:"faculty"`
	Department          bilingual_field.Bilingual  `json:"department" gorm:"embedded;embeddedPrefix:department_"`
	RequireRegistration bool                       `json:"require_registration"`
	MaxCapacity         int                        `json:"max_capacity"`
	Schedules           []schedule.Schedule        `json:"schedules" gorm:"foreignKey:event_id"`
	Location            bilingual_field.Bilingual  `json:"location" gorm:"embedded;embeddedPrefix:location_"`
	Description         *bilingual_field.Bilingual `json:"description,omitempty" gorm:"embedded;embeddedPrefix:description_"`
}

func (m Event) TableName() string {
	return "events"
}
