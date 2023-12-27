package event

import "time"

type GetAllEventResponse struct {
	events []EventDTO
}

type EventInvalidResponse struct {
	Instance string `example:"/events/:eventId" json:"instance"`
	Title    string `example:"invalid-event-id" json:"title"`
}

type EventErrorResponse struct {
	Instance string `example:"/events/:eventId"      json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type EventAllErrorResponse struct {
	Instance string `example:"/events"               json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type EventDTO struct {
	Id                  string               `example:"first-event"        json:"id"`
	Name                NameEventBilingual   `json:"name"`
	FacultyCode         int                  `json:"-"`
	Faculty             Faculty              `json:"faculty"`
	Department          DepartmentBilingual  `json:"department"`
	RequireRegistration bool                 `example:"true"               json:"require_registration"`
	MaxCapacity         int                  `example:"100"                json:"max_capacity"`
	Schedules           []Schedule           `json:"schedules"`
	Location            LocationBilingual    `json:"location"`
	Description         DescriptionBilingual `json:"description,omitempty"`
}

type NameEventBilingual struct {
	En string `example:"First Event"                    json:"en"`
	Th string `example:"อีเวนท์แรก"                     json:"th"`
}

type DepartmentBilingual struct {
	En string `example:"Computer Engineering"                                   json:"en"`
	Th string `example:"ภาควิชาคอมพิวเตอร์"                                     json:"th"`
}

type Faculty struct {
	Code string               `example:"21" json:"code"`
	Name NameFacultyBilingual `json:"name"`
}

type NameFacultyBilingual struct {
	En string `example:"Faculty of Engineering"                              json:"en"`
	Th string `example:"คณะวิศวกรรมศาสตร์"                                   json:"th"`
}

type LocationBilingual struct {
	En string `example:"SIT Building"        json:"en"`
	Th string `example:"อาคาร SIT"           json:"th"`
}

type DescriptionBilingual struct {
	En string `example:"This is the first event."                                     json:"en"`
	Th string `example:"รายละเอียดอีเวนท์แรก"                                         json:"th"`
}

type Schedule struct {
	StartsAt time.Time `example:"2021-08-01T00:00:00+07:00" json:"ends_at"`
	EndsAt   time.Time `example:"2021-08-01T00:00:00+07:00" json:"starts_at"`
}
