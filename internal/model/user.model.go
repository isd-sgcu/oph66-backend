package model

type User struct {
	ID                  int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Gender              string              `json:"gender"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
	Email               string              `gorm:"index"                    json:"email"`
	School              string              `json:"school"`
	BirthDate           string              `json:"birth_date"`
	Address             string              `json:"address"`
	FromAbroad          string              `json:"from_abroad"`
	Allergy             string              `json:"allergy"`
	MedicalCondition    string              `json:"medical_condition"`
	JoinCUReason        string              `json:"join_cu_reason"`
	NewsSource          string              `json:"news_source"`
	Status              string              `json:"status"`
	Grade               string              `json:"grade"`
	DesiredRounds       []DesiredRound      `gorm:"foreignKey:UserID"`
	InterestedFaculties []InterestedFaculty `gorm:"foreignKey:UserID"`
	RegisteredEvents    []EventRegistration `gorm:"foreignKey:UserID"`
}

func (u User) TableName() string {
	return "users"
}
