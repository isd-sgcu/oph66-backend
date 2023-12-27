package model

type Faculty struct {
	Code string    `gorm:"primaryKey"                    json:"code"`
	Name Bilingual `gorm:"embedded;embeddedPrefix:name_" json:"name"`
}

func (m Faculty) TableName() string {
	return "faculties"
}
