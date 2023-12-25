package schedule

import "time"

type schedule_period string

const (
	FIRST_MORNING    schedule_period = "20-morning"
	FIRST_AFTERNOON  schedule_period = "20-afternoon"
	SECOND_MORNING   schedule_period = "21-morning"
	SECOND_AFTERNOON schedule_period = "21-afternoon"
)

type Schedule struct {
	EventId  string          `json:"-"`
	StartsAt time.Time       `json:"ends_at"`
	EndsAt   time.Time       `json:"starts_at"`
	Period   schedule_period `json:"-"`
}

func (m Schedule) TableName() string {
	return "schedules"
}
