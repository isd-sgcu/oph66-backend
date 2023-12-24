package auth

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
	DesiredRounds       []DesiredRounds       `gorm:"foreignKey:UserID"` // One-to-many relationship
	InterestedFaculties []InterestedFaculties `gorm:"foreignKey:UserID"` // One-to-many relationship
}

type InterestedFaculties struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Section    string `json:"section"`
	UserID     uint   `gorm:"index"` // Foreign key
}

type DesiredRounds struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Round      string `json:"round"`
	UserID     uint   `gorm:"index"` // Foreign key
}

func (u *User) TableName() string {
	return "users"
}

func (d *DesiredRounds) TableName() string {
	return "desired_rounds"
}

func (i *InterestedFaculties) TableName() string {
	return "interested_faculties"
}