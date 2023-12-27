package auth

type RegisterRequestDTO struct {
	Gender              string        `json:"gender"`
	FirstName           string        `json:"first_name"`
	LastName            string        `json:"last_name"`
	School              string        `json:"school"`
	BirthDate           string        `json:"birth_date"`
	Address             string        `json:"address"`
	FromAbroad          string        `json:"from_abroad"`
	Allergy             string        `json:"allergy"`
	MedicalCondition    string        `json:"medical_condition"`
	JoinCUReason        string        `json:"join_cu_reason"`
	NewsSource          string        `json:"news_source"`
	Status              string        `json:"status"`
	Grade               string        `json:"grade"`
	DesiredRounds       []DesiredInfo `json:"desired_rounds"`
	InterestedFaculties []FacultyInfo `json:"interested_faculties"`
}

type DesiredInfo struct {
	Order uint   `json:"order"`
	Code  string `json:"code"`
}

type FacultyInfo struct {
	Order          uint   `json:"order"`
	FacultyCode    string `json:"faculty_code"`
	DepartmentCode string `json:"department_code"`
	SectionCode    string `json:"section_code"`
}

type GoogleCallbackResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	User *User `json:"user"`
}

type GetProfileResponse struct {
	User *User `json:"user"`
}
