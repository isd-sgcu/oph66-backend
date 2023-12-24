package register

import (
	"github.com/lib/pq"
)

type User struct {
	ID                  uint                  `gorm:"primaryKey;autoIncrement" json:"-"`
	Gender              string                `json:"gender"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name"`
	Email               string                `json:"email"`
	School              string                `json:"school"`
	BirthDate           string                `json:"birth_date"`
	Address             string                `json:"address"`
	FromAbroad          string                `json:"from_abroad"`
	Allergy             string                `json:"allergy"`
	MedicalCondition    string                `json:"medical_condition"`
	JoinCUReason        string                `json:"join_cu_reason"`
	NewsSource          string                `json:"news_source"`
	Status              string                `json:"status"`
	Grade               string                `json:"grade"`
	DesiredRounds       pq.Int64Array         `gorm:"type:int[]"`
	InterestedFaculties []InterestedFaculties `gorm:"foreignKey:UserID"` // One-to-many relationship
}

func (u *User) TableName() string {
	return "users"
}

type InterestedFaculties struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Section    string `json:"section"`
	UserID     uint   `gorm:"index"` // Foreign key
}
