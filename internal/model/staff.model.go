package model

import "time"

type AttendeeFacultyCheckin struct {
	Id             int        `gorm:"primaryKey,autoIncrement"`
	CreatedAt      time.Time  `gorm:"not null;autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"not null;autoUpdateTime:milli"`
	UserId         int        `gorm:"not null"`
	User           User       `gorm:"foreignKey:UserId;references:Id"`
	FacultyCode    string     `gorm:"not null"`
	Faculty        Faculty    `gorm:"foreignKey:FacultyCode;references:Code"`
	DepartmentCode string     `gorm:"not null"`
	Department     Department `gorm:"foreignKey:DepartmentCode,FacultyCode;references:Code,FacultyCode"`
}

func (AttendeeFacultyCheckin) TableName() string {
	return "attendee_faculty_chekcins"
}

type AttendeeCentralCheckin struct {
	Id        int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime:milli"`
	UserId    int       `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserId;references:Id"`
}

func (AttendeeCentralCheckin) TableName() string {
	return "attendee_central_checkins"
}
