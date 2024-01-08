package model

import "time"

type User struct {
	Id                  int                 `gorm:"primaryKey;autoIncrement"`
	CreatedAt           time.Time           `gorm:"autoCreateTime"`
	UpdatedAt           time.Time           `gorm:"autoUpdateTime:milli"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
	Email               string              `gorm:"index"`
	BirthDate           string              `json:"birth_date"`
	JoinCUReason        string              `json:"join_cu_reason"`
	Status              string              `json:"status"`
	DesiredRound        string              `gorm:""`
	Province            string              `gorm:""`
	Country             string              `gorm:""`
	EducationalLevel    string              `gorm:""`
	Allergies           string              `gorm:""`
	MedicalCondition    string              `gorm:""`
	InterestedFaculties []InterestedFaculty `gorm:"foreignKey:UserId"`
	RegisteredEvents    []EventRegistration `gorm:"foreignKey:UserId"`
	VisitingFaculties   []VisitingFaculty   `gorm:"foreignKey:UserId"`
	NewsSourceUsers     []NewsSourceUser    `gorm:"foreignKey:UserId"`
}

func (u User) TableName() string {
	return "users"
}
