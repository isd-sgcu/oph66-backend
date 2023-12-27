package model

type Round struct {
	Code string `gorm:"primaryKey" json:"code"`
	Name string `json:"name"`
}

func (r Round) TableName() string {
	return "rounds"
}
