package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func ConvertRegisterRequestDTOToUser(dto *dto.RegisterRequestDTO, email string) model.User {
	var user model.User
	user.Gender = dto.Gender
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Email = email
	user.School = dto.School
	user.BirthDate = dto.BirthDate
	user.Address = dto.Address
	user.FromAbroad = dto.FromAbroad
	user.Allergy = dto.Allergy
	user.MedicalCondition = dto.MedicalCondition
	user.JoinCUReason = dto.JoinCUReason
	user.NewsSource = dto.NewsSource
	user.Status = dto.Status
	user.Grade = dto.Grade
	user.DesiredRounds = make([]model.DesiredRound, 0, len(dto.DesiredRounds))
	user.InterestedFaculties = make([]model.InterestedFaculty, 0, len(dto.InterestedFaculties))
	user.RegisteredEvents = make([]model.EventRegistration, 0)

	for _, desiredRound := range dto.DesiredRounds {
		user.DesiredRounds = append(user.DesiredRounds, DesiredRoundDTOToModel(&desiredRound, user.Id))
	}

	for _, interestedFaculty := range dto.InterestedFaculties {
		user.InterestedFaculties = append(user.InterestedFaculties, FacultyInfoIdToInterestedFaculty(&interestedFaculty))
	}

	return user
}

func DesiredRoundDTOToModel(dto *dto.DesiredRound, userId int) model.DesiredRound {
	var desiredRound model.DesiredRound
	if dto == nil {
		return desiredRound
	}
	desiredRound.Order = dto.Order
	desiredRound.Round = model.Round(dto.Round)
	desiredRound.UserId = uint(userId)
	return desiredRound
}

func FacultyInfoIdToInterestedFaculty(dto *dto.FacultyInfoId) model.InterestedFaculty {
	var interestedFaculty model.InterestedFaculty
	interestedFaculty.Order = dto.Order
	interestedFaculty.FacultyCode = dto.FacultyCode
	interestedFaculty.DepartmentCode = dto.DepartmentCode
	interestedFaculty.SectionCode = dto.SectionCode
	return interestedFaculty
}

func UserModelToUserDTO(mUser *model.User) dto.User {
	var user dto.User
	user.Gender = mUser.Gender
	user.FirstName = mUser.FirstName
	user.LastName = mUser.LastName
	user.School = mUser.School
	user.BirthDate = mUser.BirthDate
	user.Address = mUser.Address
	user.FromAbroad = mUser.FromAbroad
	user.Allergy = mUser.Allergy
	user.MedicalCondition = mUser.MedicalCondition
	user.JoinCUReason = mUser.JoinCUReason
	user.NewsSource = mUser.NewsSource
	user.Status = mUser.Status
	user.Grade = mUser.Grade
	user.DesiredRounds = make([]dto.DesiredRound, 0, len(mUser.DesiredRounds))
	user.InterestedFaculties = make([]dto.FacultyInfo, 0, len(mUser.InterestedFaculties))

	for _, round := range mUser.DesiredRounds {
		user.DesiredRounds = append(user.DesiredRounds, DesiredRoundModelToDTO(&round))
	}

	for _, faculty := range mUser.InterestedFaculties {
		user.InterestedFaculties = append(user.InterestedFaculties, InterestedFacultyToFacultyInfo(&faculty))
	}

	for _, registeredEvent := range mUser.RegisteredEvents {
		user.RegisteredEvents = append(user.RegisteredEvents, ScheduleModelToDTO(&registeredEvent.Schedule))
	}

	return user
}

func InterestedFacultyToFacultyInfo(m *model.InterestedFaculty) dto.FacultyInfo {
	var facultyInfo dto.FacultyInfo
	if m == nil {
		return facultyInfo
	}
	facultyInfo.Department.Code = m.DepartmentCode
	facultyInfo.Department.Name.En = m.Department.Name.En
	facultyInfo.Department.Name.Th = m.Department.Name.Th
	facultyInfo.Faculty.Code = m.FacultyCode
	facultyInfo.Faculty.Name.En = m.Faculty.Name.En
	facultyInfo.Faculty.Name.Th = m.Faculty.Name.Th
	facultyInfo.Section.Code = m.SectionCode
	facultyInfo.Section.Name.En = m.Section.Name.En
	facultyInfo.Section.Name.Th = m.Section.Name.Th
	return facultyInfo
}

func DesiredRoundModelToDTO(m *model.DesiredRound) dto.DesiredRound {
	var desiredRound dto.DesiredRound
	if m == nil {
		return desiredRound
	}
	desiredRound.Order = m.Order
	desiredRound.Round = string(m.Round)
	return desiredRound
}

func ScheduleModelToDTO(m *model.Schedule) dto.Schedule {
	var registeredEvent dto.Schedule
	if m == nil {
		return registeredEvent
	}
	registeredEvent.Id = m.Id
	registeredEvent.CurrentAttendee = m.CurrentAttendee
	registeredEvent.StartsAt = m.StartsAt
	registeredEvent.EndsAt = m.EndsAt
	registeredEvent.Period = string(m.Period)
	return registeredEvent
}
