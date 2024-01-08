package model

import "time"

type NewsSource string

const (
	FACEBOOK      NewsSource = "facebook"
	INSTAGRAM     NewsSource = "instagram"
	FACULTY       NewsSource = "faculty"
	CHULA_STUDENT NewsSource = "chula-student"
	FRIEND        NewsSource = "friend"
	PARENT        NewsSource = "parent"
	SCHOOL        NewsSource = "school"
	OTHER         NewsSource = "other"
)

type NewsSourceUser struct {
	UserId     string     `gorm:"primaryKey"`
	NewsSource NewsSource `gorm:"primaryKey"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime:milli"`
}

type NewsSourceEventRegistration struct {
	UserId     int        `gorm:"primaryKey"`
	ScheduleId int        `gorm:"primaryKey"`
	NewsSource NewsSource `gorm:"primaryKey"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime:milli"`
}

func (NewsSourceUser) TableName() string {
	return "news_sources_users"
}

func (NewsSourceEventRegistration) TableName() string {
	return "news_sources_event_registrations"
}
