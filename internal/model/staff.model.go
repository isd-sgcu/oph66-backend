package model

import "time"

type AttendeeCheckin struct {
	Id             int        `gorm:"primaryKey,autoIncrement"`
	CreatedAt      time.Time  `gorm:"not null;autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"not null;autoUpdateTime:milli"`
	UserId         int        `gorm:"not null"`
	User           User       `gorm:"foreignKey:UserId;references:Id"`
	FacultyCode    string     `gorm:"not null"`
	Faculty        Faculty    `gorm:"foreignKey:FacultyCode;references:Code"`
	DepartmentCode string     `gorm:""`
	Department     Department `gorm:"foreignKey:DepartmentCode,FacultyCode;references:Code,FacultyCode"`
}
