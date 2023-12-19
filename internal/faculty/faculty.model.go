package faculty

type Faculty struct {
	Code   string `gorm:"primaryKey" json:"code"`
	NameEn string `json:"name_en"`
	NameTh string `json:"name_th"`
}

func (m Faculty) TableName() string {
	return "faculties"
}
