package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func ConvertRegisterRequestDTOToUser(dto *dto.RegisterRequestDTO, email string) model.User {
	var user model.User
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Email = email
	user.BirthDate = dto.BirthDate
	user.JoinCUReason = dto.JoinCUReason
	user.Status = dto.Status
	user.DesiredRound = dto.DesiredRound
	user.Province = dto.Province
	user.Country = dto.Country
	user.EducationalLevel = dto.EducationalLevel
	user.Allergies = dto.Allergies
	user.MedicalCondition = dto.MedicalCondition
	user.InterestedFaculties = make([]model.InterestedFaculty, 0, len(dto.InterestedFaculties))
	user.VisitingFaculties = make([]model.VisitingFaculty, 0, len(dto.VisitingFaculties))
	user.RegisteredEvents = make([]model.EventRegistration, 0)
	user.NewsSourceUsers = make([]model.NewsSourceUser, 0, len(dto.NewsSource))

	for _, visitingFaculty := range dto.VisitingFaculties {
		user.VisitingFaculties = append(user.VisitingFaculties, FacultyInfoIdToVisitingFaculty(&visitingFaculty))
	}

	for _, interestedFaculty := range dto.InterestedFaculties {
		user.InterestedFaculties = append(user.InterestedFaculties, FacultyInfoIdToInterestedFaculty(&interestedFaculty))
	}

	for _, newsSource := range dto.NewsSource {
		user.NewsSourceUsers = append(user.NewsSourceUsers, NewsSourceDTOToModel(&newsSource))
	}

	return user
}

func NewsSourceDTOToModel(dto *dto.NewsSource) model.NewsSourceUser {
	var ns model.NewsSourceUser

	ns.NewsSource = model.NewsSource(*dto)

	return ns
}

func FacultyInfoIdToInterestedFaculty(dto *dto.FacultyInfoId) model.InterestedFaculty {
	var interestedFaculty model.InterestedFaculty
	if dto == nil {
		return interestedFaculty
	}
	interestedFaculty.Order = dto.Order
	interestedFaculty.FacultyCode = dto.FacultyCode
	interestedFaculty.DepartmentCode = dto.DepartmentCode
	interestedFaculty.SectionCode = dto.SectionCode
	return interestedFaculty
}

func FacultyInfoIdToVisitingFaculty(dto *dto.FacultyInfoId) model.VisitingFaculty {
	var visitingFaculty model.VisitingFaculty
	visitingFaculty.Order = dto.Order
	visitingFaculty.FacultyCode = dto.FacultyCode
	visitingFaculty.DepartmentCode = dto.DepartmentCode
	visitingFaculty.SectionCode = dto.SectionCode
	return visitingFaculty
}

func UserModelToDTO(mUser *model.User) dto.User {
	var user dto.User
	user.Id = mUser.Id
	user.FirstName = mUser.FirstName
	user.LastName = mUser.LastName
	user.BirthDate = mUser.BirthDate
	user.JoinCUReason = mUser.JoinCUReason
	user.Status = mUser.Status
	user.DesiredRound = mUser.DesiredRound
	user.Province = mUser.Province
	user.Country = mUser.Country
	user.EducationalLevel = mUser.EducationalLevel
	user.Allergies = mUser.Allergies
	user.MedicalCondition = mUser.MedicalCondition
	user.InterestedFaculties = make([]dto.FacultyInfo, 0, len(mUser.InterestedFaculties))
	user.RegisteredEvents = make([]dto.Schedule, 0, len(mUser.RegisteredEvents))
	user.NewsSources = make([]dto.NewsSource, 0, len(mUser.NewsSourceUsers))
	user.VisitingFaculties = make([]dto.FacultyInfo, 0, len(mUser.VisitingFaculties))

	for _, faculty := range mUser.InterestedFaculties {
		user.InterestedFaculties = append(user.InterestedFaculties, InterestedFacultyToFacultyInfo(&faculty))
	}

	for _, registeredEvent := range mUser.RegisteredEvents {
		user.RegisteredEvents = append(user.RegisteredEvents, ScheduleModelToDTO(&registeredEvent.Schedule))
	}

	for _, newsSource := range mUser.NewsSourceUsers {
		user.NewsSources = append(user.NewsSources, NewsSourceModelToDTO(&newsSource))
	}

	for _, viFac := range mUser.VisitingFaculties {
		user.VisitingFaculties = append(user.VisitingFaculties, VisitingFacultyToFacultyInfo(&viFac))
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

func NewsSourceModelToDTO(m *model.NewsSourceUser) dto.NewsSource {
	var ns dto.NewsSource
	if m == nil {
		return ns
	}

	ns = dto.NewsSource(m.NewsSource)

	return ns
}

func VisitingFacultyToFacultyInfo(m *model.VisitingFaculty) dto.FacultyInfo {
	var fi dto.FacultyInfo
	if m == nil {
		return fi
	}
	fi.Department.Code = m.DepartmentCode
	fi.Department.Name.En = m.Department.Name.En
	fi.Department.Name.Th = m.Department.Name.Th
	fi.Faculty.Code = m.FacultyCode
	fi.Faculty.Name.En = m.Faculty.Name.En
	fi.Faculty.Name.Th = m.Faculty.Name.Th
	fi.Section.Code = m.SectionCode
	fi.Section.Name.En = m.Section.Name.En
	fi.Section.Name.Th = m.Section.Name.Th
	return fi
}
