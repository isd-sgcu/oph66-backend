package model

type EventRegistration struct {
	User       User     `gorm:"foreignKey:UserID"`
	UserID     int      `gorm:""`
	Schedule   Schedule `gorm:"foreignKey:ScheduleID"`
	ScheduleID int      `gorm:""`
}
