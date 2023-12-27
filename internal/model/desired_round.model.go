package model

type DesiredRound struct {
	UserID    uint   `gorm:"primaryKey;index"     json:"-"`
	Order     uint   `gorm:"primaryKey"           json:"order"`
	Round     Round  `gorm:"foreignKey:RoundCode" json:"round"`
	RoundCode string `json:"-"`
}

func (d DesiredRound) TableName() string {
	return "desired_rounds"
}
