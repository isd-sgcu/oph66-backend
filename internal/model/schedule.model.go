package model

import "time"

type SchedulePeriod string

const (
	FIRST_MORNING    SchedulePeriod = "20-morning"
	FIRST_AFTERNOON  SchedulePeriod = "20-afternoon"
	SECOND_MORNING   SchedulePeriod = "21-morning"
	SECOND_AFTERNOON SchedulePeriod = "21-afternoon"
)

type Schedule struct {
	ID       int            `gorm:"primaryKey"`
	EventId  int            `json:"-"`
	StartsAt time.Time      `json:"ends_at"`
	EndsAt   time.Time      `json:"starts_at"`
	Period   SchedulePeriod `json:"-"`
}

func (m Schedule) TableName() string {
	return "schedules"
}
