package evtreg

import (
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func CombineSchedultToEventRegistration(scheduleId int, userId int, newsSources []dto.NewsSource) model.EventRegistration {
	var evtreg model.EventRegistration
	evtreg.UserId = userId
	evtreg.ScheduleId = scheduleId
	evtreg.NewsSourceRegistrations = make([]model.NewsSourceEventRegistration, 0, len(newsSources))

	for _, newsSource := range newsSources {
		evtreg.NewsSourceRegistrations = append(evtreg.NewsSourceRegistrations, ConvertNewsSourceRegistrationDTOToModel(&newsSource, userId, scheduleId))
	}

	return evtreg
}

func ConvertNewsSourceRegistrationDTOToModel(dto *dto.NewsSource, userId int, scheduleId int) model.NewsSourceEventRegistration {
	var ns model.NewsSourceEventRegistration
	if dto == nil {
		return ns
	}
	ns.UserId = userId
	ns.ScheduleId = scheduleId
	ns.NewsSource = model.NewsSource(*dto)

	return ns
}
