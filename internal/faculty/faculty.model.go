package faculty

import "github.com/isd-sgcu/oph66-backend/internal/bilingual_field"

type Faculty struct {
	Code string                    `gorm:"primaryKey"                    json:"code"`
	Name bilingual_field.Bilingual `gorm:"embedded;embeddedPrefix:name_" json:"name"`
}

func (m Faculty) TableName() string {
	return "faculties"
}
