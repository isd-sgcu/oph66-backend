package model

type Department struct {
	Code        string    `gorm:"primaryKey"                    json:"code"`
	Name        Bilingual `gorm:"embedded;embeddedPrefix:name_" json:"name"`
	FacultyCode string    `gorm:"primaryKey"`
	Faculty     Faculty   `gorm:"foreignKey:FacultyCode"`
}

func (d Department) TableName() string {
	return "departments"
}
