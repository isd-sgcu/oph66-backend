package event

import (
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func EventModelToDTO(m *model.Event) dto.Event {
	var event dto.Event
	event.Id = m.Id
	event.Name = dto.BilingualModelToDTO(&m.Name)
	event.Faculty.Code = m.FacultyCode
	event.Faculty.Name = dto.BilingualModelToDTO(&m.Faculty.Name)
	event.Department.Code = m.DepartmentCode
	event.Department.Name = dto.BilingualModelToDTO(&m.Department.Name)
	event.RequireRegistration = m.RequireRegistration
	event.MaxCapacity = m.MaxCapacity
	event.Schedules = make([]dto.Schedule, 0, len(m.Schedules))
	for _, schedule := range m.Schedules {
		event.Schedules = append(event.Schedules, ScheduleModelToDTO(&schedule))
	}
	event.Location = dto.BilingualModelToDTO(&m.Location)
	event.Description = dto.BilingualModelToDTO(m.Description)
	return event
}

func ScheduleModelToDTO(m *model.Schedule) dto.Schedule {
	var schedule dto.Schedule
	schedule.Id = m.Id
	schedule.EndsAt = m.EndsAt
	schedule.StartsAt = m.StartsAt
	schedule.CurrentAttendee = m.CurrentAttendee
	schedule.Period = string(m.Period)
	return schedule
}
