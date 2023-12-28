package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func ConvertRegisterRequestDTOToUser(dto *RegisterRequestDTO, email string) *model.User {
	user := &model.User{}
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
	user.DesiredRounds = make([]model.DesiredRound, len(dto.DesiredRounds))
	user.InterestedFaculties = make([]model.InterestedFaculty, len(dto.InterestedFaculties))
	user.RegisteredEvents = make([]model.EventRegistration, 0)

	for i, desiredRound := range dto.DesiredRounds {
		ConvertDesiredInfoToDesiredRound(&desiredRound, user, &user.DesiredRounds[i])
	}

	for i, interestedFaculty := range dto.InterestedFaculties {
		ConvertFacultyInfoToInterestedFaculty(&interestedFaculty, user, &user.InterestedFaculties[i])
	}

	return user
}

func ConvertDesiredInfoToDesiredRound(dto *DesiredRound, user *model.User, desiredRound *model.DesiredRound) {
}

func ConvertFacultyInfoToInterestedFaculty(dto *FacultyInfoId, user *model.User, interestedFaculty *model.InterestedFaculty) {
	interestedFaculty.Order = dto.Order
	interestedFaculty.FacultyCode = dto.FacultyCode
	interestedFaculty.DepartmentCode = dto.DepartmentCode
	interestedFaculty.SectionCode = dto.SectionCode
}

func ConvertUserModelToUserDTO(mUser *model.User) User {
	user := User{
		Gender:              mUser.Gender,
		FirstName:           mUser.FirstName,
		LastName:            mUser.LastName,
		School:              mUser.School,
		BirthDate:           mUser.BirthDate,
		Address:             mUser.Address,
		FromAbroad:          mUser.FromAbroad,
		Allergy:             mUser.Allergy,
		MedicalCondition:    mUser.MedicalCondition,
		JoinCUReason:        mUser.JoinCUReason,
		NewsSource:          mUser.NewsSource,
		Status:              mUser.Status,
		Grade:               mUser.Grade,
		DesiredRounds:       make([]DesiredRound, 0, len(mUser.DesiredRounds)),
		InterestedFaculties: make([]FacultyInfo, 0, len(mUser.InterestedFaculties)),
	}

	for _, round := range mUser.DesiredRounds {
		user.DesiredRounds = append(user.DesiredRounds, ConvertDesiredRoundModelToDTO(&round))
	}

	for _, faculty := range mUser.InterestedFaculties {
		user.InterestedFaculties = append(user.InterestedFaculties, ConvertInterestedFacultyToFacultyInfo(&faculty))
	}

	return user
}

func ConvertInterestedFacultyToFacultyInfo(m *model.InterestedFaculty) FacultyInfo {
	var facultyInfo FacultyInfo
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

func ConvertDesiredRoundModelToDTO(m *model.DesiredRound) DesiredRound {
	var dr DesiredRound
	dr.Order = m.Order
	dr.Round = string(m.Round)
	return dr
}
