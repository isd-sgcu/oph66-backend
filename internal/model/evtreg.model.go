package model

type EventRegistration struct {
	User       User     `gorm:"foreignKey:UserId"`
	UserID     int      `gorm:""`
	Schedule   Schedule `gorm:"foreignKey:ScheduleId"`
	ScheduleId int      `gorm:""`
}
