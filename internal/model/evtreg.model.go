package model

type EventRegistration struct {
	User       User     `gorm:"foreignKey:UserId"`
	UserId     int      `gorm:""`
	Schedule   Schedule `gorm:"foreignKey:ScheduleId"`
	ScheduleId int      `gorm:""`
}
