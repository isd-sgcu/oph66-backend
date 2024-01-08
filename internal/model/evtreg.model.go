package model

import "time"

type EventRegistration struct {
	User                    User                          `gorm:"foreignKey:UserId"`
	UserId                  int                           `gorm:""`
	Schedule                Schedule                      `gorm:"foreignKey:ScheduleId"`
	ScheduleId              int                           `gorm:""`
	CreatedAt               time.Time                     `gorm:"autoCreateTime"`
	UpdatedAt               time.Time                     `gorm:"autoUpdateTime:milli"`
	NewsSourceRegistrations []NewsSourceEventRegistration `gorm:"foreignKey:UserId,ScheduleId;references:UserId,ScheduleId"`
}
