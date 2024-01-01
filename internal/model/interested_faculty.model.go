package model

import "time"

type InterestedFaculty struct {
	UserId         uint       `gorm:"primaryKey;index"                                  json:"-"`
	Order          uint       `gorm:"primaryKey"                                        json:"order"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime:milli"`
	Faculty        Faculty    `gorm:"foreignKey:FacultyCode"                            json:"faculty"`
	FacultyCode    string     `gorm:"not null"                                          json:"-"`
	Department     Department `gorm:"foreignKey:DepartmentCode,FacultyCode"             json:"department"`
	DepartmentCode string     `gorm:"not null"                                          json:"-"`
	Section        Section    `gorm:"foreignKey:SectionCode,DepartmentCode,FacultyCode" json:"section"`
	SectionCode    string     `gorm:"not null"                                          json:"-"`
}

func (i InterestedFaculty) TableName() string {
	return "interested_faculties"
}
