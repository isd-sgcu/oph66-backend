package model

type Section struct {
	Code           string     `gorm:"primaryKey"                            json:"code"`
	Department     Department `gorm:"foreignKey:FacultyCode,DepartmentCode"`
	DepartmentCode string     `gorm:"primaryKey;not null"`
	Faculty        Faculty    `gorm:"foreignKey:FacultyCode"`
	FacultyCode    string     `gorm:"primaryKey;not null"`
	Name           Bilingual  `gorm:"embedded;embeddedPrefix:name_"         json:"name"`
}

func (u Section) TableName() string {
	return "sections"
}
