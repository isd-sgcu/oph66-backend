package event

import "time"

type Event struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Description string
	StartTime   time.Time
	Location    string
	MaxCapacity int
	Department  string
}

func (m Event) TableName() string {
	return "events"
}
