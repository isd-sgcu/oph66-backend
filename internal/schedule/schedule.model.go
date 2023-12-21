package schedule

import "time"

type Schedule struct {
	EventId  string    `json:"-"`
	StartsAt time.Time `json:"ends_at"`
	EndsAt   time.Time `json:"starts_at"`
}

func (m Schedule) TableName() string {
	return "schedules"
}
