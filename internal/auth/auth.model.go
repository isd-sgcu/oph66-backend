package auth

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
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Order      uint   `json:"order"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Section    string `json:"section"`
	UserID     uint   `gorm:"index"`
}

type DesiredRound struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Order  uint   `json:"order"`
	Round  string `json:"round"`
	UserID uint   `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

func (d *DesiredRound) TableName() string {
	return "desired_round"
}

func (i *InterestedFaculty) TableName() string {
	return "interested_faculty"
}
