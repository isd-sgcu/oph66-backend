package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/faculty"
)

type User struct {
	ID                  uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	Gender              string              `json:"gender"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
	Email               string              `gorm:"index"                    json:"email"`
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
	Order          uint            `gorm:"primaryKey"                                json:"order"`
	Faculty        faculty.Faculty `gorm:"foreignKey:FacultyCode;references:Code"    json:"faculty"`
	Department     Department      `gorm:"foreignKey:DepartmentCode;references:Code" json:"department"`
	Section        Section         `gorm:"foreignKey:SectionCode;references:Code"    json:"section"`
	FacultyCode    string          `json:"-"`
	DepartmentCode string          `json:"-"`
	SectionCode    string          `json:"-"`
	UserID         uint            `gorm:"primaryKey;index"                          json:"-"`
}

type DesiredRound struct {
	Order     uint   `gorm:"primaryKey"                           json:"order"`
	Round     Round  `gorm:"foreignKey:RoundCode;references:Code" json:"round"`
	UserID    uint   `gorm:"primaryKey;index"                     json:"-"`
	RoundCode string `json:"-"`
}

type Department struct {
	Code string `gorm:"primaryKey" json:"code"`
	Name string `json:"name"`
}

type Section struct {
	Code string `gorm:"primaryKey" json:"code"`
	Name string `json:"name"`
}

type Round struct {
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
