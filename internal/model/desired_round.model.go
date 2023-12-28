package model

type DesiredRound struct {
	UserId uint  `gorm:"primaryKey;index" json:"-"`
	Order  uint  `gorm:"primaryKey"       json:"order"`
	Round  Round `gorm:""                 json:"round"`
}

func (d DesiredRound) TableName() string {
	return "desired_rounds"
}
