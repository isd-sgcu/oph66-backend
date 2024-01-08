package dto

type RegisterRequestDTO struct {
	FirstName           string          `example:"John"                                                                json:"first_name"`
	LastName            string          `example:"Doe"                                                                 json:"last_name"`
	BirthDate           string          `example:"1990-01-01"                                                          json:"birth_date"`
	JoinCUReason        string          `example:"Interested in the programs offered"                                  json:"join_cu_reason"`
	Status              string          `example:"student"                                                             json:"status"`
	DesiredRound        string          `example:"3"                                                                   json:"desired_round"`
	Country             string          `example:"Japan"                                                               json:"country"`
	Province            string          `example:"Tokyo"                                                               json:"province"`
	EducationalLevel    string          `example:"Ph.D."                                                               json:"educational_level"`
	Allergies           string          `example:"Dog"                                                                 json:"allergies"`
	MedicalCondition    string          `example:"Dog"                                                                 json:"medical_condition"`
	InterestedFaculties []FacultyInfoId `json:"interested_faculties"`
	VisitingFaculties   []FacultyInfoId `json:"visiting_faculties"`
	NewsSource          []NewsSource    `example:"facebook,instagram,faculty,chula-student,friend,parent,school,other" json:"news_sources"`
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
	FirstName           string        `example:"John"                               json:"first_name"`
	LastName            string        `example:"Doe"                                json:"last_name"`
	BirthDate           string        `example:"1990-01-01"                         json:"birth_date"`
	JoinCUReason        string        `example:"Interested in the programs offered" json:"join_cu_reason"`
	Status              string        `example:"student"                            json:"status"`
	Country             string        `example:"Japan"                              json:"country"`
	Province            string        `example:"Austin"                             json:"province"`
	DesiredRound        string        `json:"desired_round"`
	EducationalLevel    string        `example:"Ph.D."                              json:"educational_level"`
	Allergies           string        `example:"Dog"                                json:"allergies"`
	MedicalCondition    string        `example:"Dog"                                json:"medical_condition"`
	NewsSources         []NewsSource  `example:"facebook,instagram"                 json:"news_sources"`
	InterestedFaculties []FacultyInfo `json:"interested_faculties"`
	RegisteredEvents    []Schedule    `json:"registered_events"`
	VisitingFaculties   []FacultyInfo `json:"visiting_faculties"`
}

type NewsSource string

type FacultyInfo struct {
	Faculty struct {
		Name BilingualField `json:"name"`
		Code string         `json:"code"`
	} `json:"faculty"`
	Department struct {
		Name BilingualField `json:"name"`
		Code string         `json:"code"`
	} `json:"department"`
	Section struct {
		Name BilingualField `json:"name"`
		Code string         `json:"code"`
	} `json:"section"`
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
