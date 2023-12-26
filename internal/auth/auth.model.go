package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/faculty"
)

type User struct {
	ID                  uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Gender              string              `json:"gender"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
	Email               string              `json:"email"`
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
}

type InterestedFaculty struct {
	ID             uint            `gorm:"primaryKey;autoIncrement"                  json:"id"`
	Order          uint            `json:"order"`
	Faculty        faculty.Faculty `gorm:"foreignKey:FacultyCode;references:Code"    json:"faculty"`
	Department     department      `gorm:"foreignKey:DepartmentCode;references:Code" json:"department"`
	Section        section         `gorm:"foreignKey:SectionCode;references:Code"    json:"section"`
	FacultyCode    string          `json:"-"`
	DepartmentCode string          `json:"-"`
	SectionCode    string          `json:"-"`
	UserID         uint            `gorm:"index"`
}

type DesiredRound struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"             json:"id"`
	Order     uint   `json:"order"`
	Round     round  `gorm:"foreignKey:RoundCode;references:Code" json:"round"`
	UserID    uint   `gorm:"index"`
	RoundCode string `json:"-"`
}

type department struct {
	Code    string `gorm:"primaryKey"      json:"code"`
	Name    string `json:"name"`
}

type section struct {
	Code       string `gorm:"primaryKey"      json:"code"`
	Name       string `json:"name"`
}

type round struct {
	Code string `gorm:"primaryKey" json:"code"`
	Name string `json:"name"`
}

func (u *User) TableName() string {
	return "users"
}

func (d *DesiredRound) TableName() string {
	return "desired_rounds"
}

func (i *InterestedFaculty) TableName() string {
	return "interested_faculties"
}
