package model

type Event struct {
	Id                  string     `gorm:"primaryKey"                           json:"id"`
	Name                Bilingual  `gorm:"embedded;embeddedPrefix:name_"        json:"name"`
	FacultyCode         int        `gorm:"references:Code"                      json:"-"`
	Faculty             Faculty    `json:"faculty"`
	Department          Bilingual  `gorm:"embedded;embeddedPrefix:department_"  json:"department"`
	RequireRegistration bool       `json:"require_registration"`
	MaxCapacity         int        `json:"max_capacity"`
	Schedules           []Schedule `gorm:"foreignKey:event_id"                  json:"schedules"`
	Location            Bilingual  `gorm:"embedded;embeddedPrefix:location_"    json:"location"`
	Description         *Bilingual `gorm:"embedded;embeddedPrefix:description_" json:"description,omitempty"`
}

func (m Event) TableName() string {
	return "events"
}
