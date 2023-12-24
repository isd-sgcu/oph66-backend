package event

import "time"

type EventAll []Event

type EventInvalidResponse struct {
	Instance string `json:"instance" example:"/events/:eventId"`
	Title    string `json:"title" example:"invalid-event-id"`
}

type EventErrorResponse struct {
	Instance string `json:"instance" example:"/events/:eventId"`
	Title    string `json:"title" example:"internal-server-error"`
}

type EventAllErrorResponse struct {
	Instance string `json:"instance" example:"/events"`
	Title    string `json:"title" example:"internal-server-error"`
}

type EventDTO struct {
	Id                  string               `json:"id" example:"first-event"`
	Name                NameEventBilingual   `json:"name"`
	FacultyCode         int                  `json:"-"`
	Faculty             Faculty              `json:"faculty"`
	Department          DepartmentBilingual  `json:"department"`
	RequireRegistration bool                 `json:"require_registration" example:"true"`
	MaxCapacity         int                  `json:"max_capacity" example:"100"`
	Schedules           []Schedule           `json:"schedules"`
	Location            LocationBilingual    `json:"location"`
	Description         DescriptionBilingual `json:"description,omitempty"`
}

type NameEventBilingual struct {
	En string `json:"en" example:"First Event"`
	Th string `json:"th" example:"อีเวนท์แรก"`
}

type DepartmentBilingual struct {
	En string `json:"en" example:"Computer Engineering"`
	Th string `json:"th" example:"ภาควิชาคอมพิวเตอร์"`
}

type Faculty struct {
	Code string               `json:"code" example:"21"`
	Name NameFacultyBilingual `json:"name"`
}

type NameFacultyBilingual struct {
	En string `json:"en" example:"Faculty of Engineering"`
	Th string `json:"th" example:"คณะวิศวกรรมศาสตร์"`
}

type LocationBilingual struct {
	En string `json:"en" example:"SIT Building"`
	Th string `json:"th" example:"อาคาร SIT"`
}

type DescriptionBilingual struct {
	En string `json:"en" example:"This is the first event."`
	Th string `json:"th" example:"รายละเอียดอีเวนท์แรก"`
}

type Schedule struct {
	StartsAt time.Time `json:"ends_at" example:"2021-08-01T00:00:00+07:00"`
	EndsAt   time.Time `json:"starts_at" example:"2021-08-01T00:00:00+07:00"`
}
