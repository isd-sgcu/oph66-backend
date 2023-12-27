package auth

import "github.com/isd-sgcu/oph66-backend/internal/model"

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
	User *model.User `json:"user"`
}

type GetProfileResponse struct {
	User *model.User `json:"user"`
}

type MockUser struct {
	Gender              string                  `example:"male"                               json:"gender"`
	FirstName           string                  `example:"John"                               json:"first_name"`
	LastName            string                  `example:"Doe"                                json:"last_name"`
	School              string                  `example:"CU"                                 json:"school"`
	BirthDate           string                  `example:"1990-01-01"                         json:"birth_date"`
	Address             string                  `example:"Bangkok"                            json:"address"`
	FromAbroad          string                  `example:"no"                                 json:"from_abroad"`
	Allergy             string                  `example:"None"                               json:"allergy"`
	MedicalCondition    string                  `example:"None"                               json:"medical_condition"`
	JoinCUReason        string                  `example:"Interested in the programs offered" json:"join_cu_reason"`
	NewsSource          string                  `example:"Facebook"                           json:"news_source"`
	Status              string                  `example:"student"                            json:"status"`
	Grade               string                  `example:"undergraduate"                      json:"grade"`
	DesiredRounds       []MockDesiredRound      `json:"desired_rounds"`
	InterestedFaculties []MockInterestedFaculty `json:"interested_faculties"`
}

type MockDesiredRound struct {
	Order uint   `example:"1" json:"order"`
	Code  string `example:"1" json:"code"`
}

type MockInterestedFaculty struct {
	Order uint   `example:"1" json:"order"`
	Code  string `example:"1" json:"code"`
}

type CallbackResponse struct {
	Token string `example:"gbxnZjiHVzb_4mDQTQNiJdrZFOCactWXkZvZOxS2_qZsy7vAQY7uA2RFIHe2JABoEjhT0Y3KlOJuOEvE2YJMLrJDagwhpAITGex" json:"token"`
}

type CallbackErrorResponse struct {
	Instance string `example:"/auth/callback"        json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type CallbackInvalidResponse struct {
	Instance string `example:"/auth/callback" json:"instance"`
	Title    string `example:"bad-request"    json:"title"`
}

type MockRegisterResponse struct {
	MockUser MockUser `json:"user"`
}

type RegisterErrorResponse struct {
	Instance string `example:"/auth/register"        json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type RegisterInvalidResponse struct {
	Instance string `example:"/auth/register" json:"instance"`
	Title    string `example:"bad-request"    json:"title"`
}

type RegisterUnauthorized struct {
	Instance string `example:"/auth/register" json:"instance"`
	Title    string `example:"unauthorized"   json:"title"`
}

type RegisterInvalidToken struct {
	Instance string `example:"/auth/register" json:"instance"`
	Title    string `example:"invalid-token"  json:"title"`
}

type MockGetProfileResponse struct {
	MockUser MockUser `json:"user"`
}

type GetProfileErrorResponse struct {
	Instance string `example:"/auth/me"              json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type GetProfileUnauthorized struct {
	Instance string `example:"/auth/me"     json:"instance"`
	Title    string `example:"unauthorized" json:"title"`
}

type GetProfileUserNotFound struct {
	Instance string `example:"/auth/me"       json:"instance"`
	Title    string `example:"user-not-found" json:"title"`
}
