package model

type Section struct {
	Code           string     `gorm:"primaryKey" json:"code"`
	Faculty        Faculty    `gorm:"primaryKey;foreignKey:FacultyCode"`
	FacultyCode    string     `gorm:"not null"`
	Department     Department `gorm:"primaryKey;foreignKey:FacultyCode,DepartmentCode"`
	DepartmentCode string     `gorm:"not null"`
	Name           string     `json:"name"`
}

func (u Section) TableName() string {
	return "sections"
}
