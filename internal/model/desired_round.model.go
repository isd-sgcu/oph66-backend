package model

import "time"

type DesiredRound struct {
	UserId    uint      `gorm:"primaryKey;index" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	Order     uint      `gorm:"primaryKey"       json:"order"`
	Round     Round     `gorm:""                 json:"round"`
}

func (d DesiredRound) TableName() string {
	return "desired_rounds"
}
