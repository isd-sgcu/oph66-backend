package faculty

import "github.com/isd-sgcu/oph66-backend/internal/bilingual_field"

type Faculty struct {
	Code string                    `json:"code" gorm:"primaryKey"`
	Name bilingual_field.Bilingual `json:"name" gorm:"embedded;embeddedPrefix:name_"`
}

func (m Faculty) TableName() string {
	return "faculties"
}
