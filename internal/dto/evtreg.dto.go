package dto

type EventRegistrationDTO struct {
	NewsSource []NewsSource `example:"facebook,instagram,faculty,chula-student,friend,parent,school,other" json:"news_sources"`
}
