package model

type Department struct {
	Code        string  `gorm:"primaryKey"             json:"code"`
	Name        string  `json:"name"`
	FacultyCode string  `gorm:"primaryKey"`
	Faculty     Faculty `gorm:"foreignKey:FacultyCode"`
}

func (d Department) TableName() string {
	return "departments"
}
