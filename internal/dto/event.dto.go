package dto

import (
	"time"
)

type GetAllEventResponse struct {
	Events []Event `json:"events"`
}

type GetEventByIdResponse struct {
	Event Event `json:"event"`
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

type Event struct {
	Id                  string         `example:"first-event"        json:"id"`
	Name                BilingualField `json:"name"`
	Faculty             Faculty        `json:"faculty"`
	Department          Department     `json:"department"`
	RequireRegistration bool           `example:"true"               json:"require_registration"`
	MaxCapacity         int            `example:"100"                json:"max_capacity"`
	Schedules           []Schedule     `json:"schedules"`
	Location            BilingualField `json:"location"`
	Description         BilingualField `json:"description,omitempty"`
}

type Faculty struct {
	Code string         `example:"21" json:"code"`
	Name BilingualField `json:"name"`
}

type Department struct {
	Code string         `example:"21" json:"code"`
	Name BilingualField `json:"name"`
}

type Schedule struct {
	Id              int       `example:"5" json:"id"`
	CurrentAttendee int       `example:"83" json:"current_attendee"`
	StartsAt        time.Time `example:"2021-08-01T00:00:00+07:00" json:"ends_at"`
	EndsAt          time.Time `example:"2021-08-01T00:00:00+07:00" json:"starts_at"`
	Period          string    `example:"20-morning" json:"period"`
}
