package event

import (
	"github.com/isd-sgcu/oph66-backend/internal/bilingual_field"
	"github.com/isd-sgcu/oph66-backend/internal/faculty"
	"github.com/isd-sgcu/oph66-backend/internal/schedule"
)

type Event struct {
	Id                  string                     `gorm:"primaryKey"                           json:"id"`
	Name                bilingual_field.Bilingual  `gorm:"embedded;embeddedPrefix:name_"        json:"name"`
	FacultyCode         int                        `gorm:"references:Code"                      json:"-"`
	Faculty             faculty.Faculty            `json:"faculty"`
	Department          bilingual_field.Bilingual  `gorm:"embedded;embeddedPrefix:department_"  json:"department"`
	RequireRegistration bool                       `json:"require_registration"`
	MaxCapacity         int                        `json:"max_capacity"`
	Schedules           []schedule.Schedule        `gorm:"foreignKey:event_id"                  json:"schedules"`
	Location            bilingual_field.Bilingual  `gorm:"embedded;embeddedPrefix:location_"    json:"location"`
	Description         *bilingual_field.Bilingual `gorm:"embedded;embeddedPrefix:description_" json:"description,omitempty"`
}

func (m Event) TableName() string {
	return "events"
}
