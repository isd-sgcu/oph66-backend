package auth

type RegisterRequestDTO struct {
	Gender              string          `example:"male"                               json:"gender"`
	FirstName           string          `example:"John"                               json:"first_name"`
	LastName            string          `example:"Doe"                                json:"last_name"`
	School              string          `example:"CU"                                 json:"school"`
	BirthDate           string          `example:"1990-01-01"                         json:"birth_date"`
	Address             string          `example:"Bangkok"                            json:"address"`
	FromAbroad          string          `example:"no"                                 json:"from_abroad"`
	Allergy             string          `example:"None"                               json:"allergy"`
	MedicalCondition    string          `example:"None"                               json:"medical_condition"`
	JoinCUReason        string          `example:"Interested in the programs offered" json:"join_cu_reason"`
	NewsSource          string          `example:"Facebook"                           json:"news_source"`
	Status              string          `example:"student"                            json:"status"`
	Grade               string          `example:"undergraduate"                      json:"grade"`
	DesiredRounds       []DesiredRound  `json:"desired_rounds"`
	InterestedFaculties []FacultyInfoId `json:"interested_faculties"`
}

type DesiredRound struct {
	Order uint   `example:"1" json:"order"`
	Round string `example:"1" json:"round"`
}

type FacultyInfoId struct {
	Order          uint   `example:"1"  json:"order"`
	FacultyCode    string `example:"21" json:"faculty_code"`
	DepartmentCode string `example:"10" json:"department_code"`
	SectionCode    string `example:"-"  json:"section_code"`
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

type User struct {
	Gender              string         `example:"male"                               json:"gender"`
	FirstName           string         `example:"John"                               json:"first_name"`
	LastName            string         `example:"Doe"                                json:"last_name"`
	School              string         `example:"CU"                                 json:"school"`
	BirthDate           string         `example:"1990-01-01"                         json:"birth_date"`
	Address             string         `example:"Bangkok"                            json:"address"`
	FromAbroad          string         `example:"no"                                 json:"from_abroad"`
	Allergy             string         `example:"None"                               json:"allergy"`
	MedicalCondition    string         `example:"None"                               json:"medical_condition"`
	JoinCUReason        string         `example:"Interested in the programs offered" json:"join_cu_reason"`
	NewsSource          string         `example:"Facebook"                           json:"news_source"`
	Status              string         `example:"student"                            json:"status"`
	Grade               string         `example:"undergraduate"                      json:"grade"`
	DesiredRounds       []DesiredRound `json:"desired_rounds"`
	InterestedFaculties []FacultyInfo  `json:"interested_faculties"`
}

type FacultyInfo struct {
	Faculty struct {
		Name BillingualName `json:"name"`
		Code string         `json:"code"`
	} `json:"faculty"`
	Department struct {
		Name BillingualName `json:"name"`
		Code string         `json:"code"`
	} `json:"department"`
	Section struct {
		Name BillingualName `json:"name"`
		Code string         `json:"code"`
	} `json:"section"`
}

type BillingualName struct {
	Th string `json:"th"`
	En string `json:"en"`
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
