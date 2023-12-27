package model

type Round struct {
	RoundNo int `gorm:"primaryKey"`
}

func (r Round) TableName() string {
	return "rounds"
}
