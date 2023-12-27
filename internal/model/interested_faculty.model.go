package model

type InterestedFaculty struct {
	UserID         uint       `gorm:"primaryKey;index"                      json:"-"`
	Order          uint       `gorm:"primaryKey"                            json:"order"`
	Faculty        Faculty    `gorm:"foreignKey:FacultyCode"                json:"faculty"`
	FacultyCode    string     `gorm:"not null"`
	Department     Department `gorm:"foreignKey:FacultyCode,DepartmentCode" json:"department"`
	DepartmentCode string     `gorm:"not null"`
}

func (i InterestedFaculty) TableName() string {
	return "interested_faculties"
}
