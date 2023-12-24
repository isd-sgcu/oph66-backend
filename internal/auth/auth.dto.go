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

type RegisterDTO struct {
	Gender              string              `json:"gender"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
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
	DesiredRounds       []DesiredRound      `json:"desired_rounds"`
	InterestedFaculties []InterestedFaculty `json:"interested_faculties"`
}

type MockUser struct {
	Gender              string                  `json:"gender" example:"male"`
	FirstName           string                  `json:"first_name" example:"John"`
	LastName            string                  `json:"last_name" example:"Doe"`
	School              string                  `json:"school" example:"CU"`
	BirthDate           string                  `json:"birth_date" example:"1990-01-01"`
	Address             string                  `json:"address" example:"Bangkok"`
	FromAbroad          string                  `json:"from_abroad" example:"no"`
	Allergy             string                  `json:"allergy" example:"None"`
	MedicalCondition    string                  `json:"medical_condition" example:"None"`
	JoinCUReason        string                  `json:"join_cu_reason" example:"Interested in the programs offered"`
	NewsSource          string                  `json:"news_source" example:"Facebook"`
	Status              string                  `json:"status" example:"student"`
	Grade               string                  `json:"grade" example:"undergraduate"`
	DesiredRounds       []MockDesiredRound      `json:"desired_rounds"`
	InterestedFaculties []MockInterestedFaculty `json:"interested_faculties"`
}

type MockDesiredRound struct {
	Order uint   `json:"order" example:"1"`
	Code  string `json:"code" example:"1"`
}

type MockInterestedFaculty struct {
	Order uint   `json:"order" example:"1"`
	Code  string `json:"code" example:"1"`
}

type CallbackResponse struct {
	Token string `json:"token" example:"gbxnZjiHVzb_4mDQTQNiJdrZFOCactWXkZvZOxS2_qZsy7vAQY7uA2RFIHe2JABoEjhT0Y3KlOJuOEvE2YJMLrJDagwhpAITGex"`
}

type CallbackErrorResponse struct {
	Instance string `json:"instance" example:"/auth/callback"`
	Title    string `json:"title" example:"internal-server-error"`
}

type CallbackInvalidResponse struct {
	Instance string `json:"instance" example:"/auth/callback"`
	Title    string `json:"title" example:"bad-request"`
}

type MockRegisterResponse struct {
	MockUser MockUser `json:"user"`
}

type RegisterErrorResponse struct {
	Instance string `json:"instance" example:"/auth/register"`
	Title    string `json:"title" example:"internal-server-error"`
}

type RegisterInvalidResponse struct {
	Instance string `json:"instance" example:"/auth/register"`
	Title    string `json:"title" example:"bad-request"`
}

type RegisterUnauthorized struct {
	Instance string `json:"instance" example:"/auth/register"`
	Title    string `json:"title" example:"unauthorized"`
}

type RegisterInvalidToken struct {
	Instance string `json:"instance" example:"/auth/register"`
	Title    string `json:"title" example:"invalid-token"`
}

type MockGetProfileResponse struct {
	MockUser MockUser `json:"user"`
}

type GetProfileErrorResponse struct {
	Instance string `json:"instance" example:"/auth/me"`
	Title    string `json:"title" example:"internal-server-error"`
}

type GetProfileUnauthorized struct {
	Instance string `json:"instance" example:"/auth/me"`
	Title    string `json:"title" example:"unauthorized"`
}

type GetProfileUserNotFound struct {
	Instance string `json:"instance" example:"/auth/me"`
	Title    string `json:"title" example:"user-not-found"`
}
