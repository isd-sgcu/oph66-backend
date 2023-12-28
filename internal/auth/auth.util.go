package auth

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func ConvertRegisterRequestDTOToUser(user *model.User, dto *RegisterRequestDTO, email string) {
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
		DesiredRoundDTOToModel(&user.DesiredRounds[i], &desiredRound, user)
	}

	for i, interestedFaculty := range dto.InterestedFaculties {
		FacultyInfoToInterestedFaculty(&user.InterestedFaculties[i], &interestedFaculty, user)
	}
}

func DesiredRoundDTOToModel(desiredRound *model.DesiredRound, dto *DesiredRound, user *model.User) {
	desiredRound.Order = dto.Order
	desiredRound.Round = model.Round(dto.Round)
	desiredRound.UserID = uint(user.ID)
}

func FacultyInfoToInterestedFaculty(interestedFaculty *model.InterestedFaculty, dto *FacultyInfoId, user *model.User) {
	interestedFaculty.Order = dto.Order
	interestedFaculty.FacultyCode = dto.FacultyCode
	interestedFaculty.DepartmentCode = dto.DepartmentCode
	interestedFaculty.SectionCode = dto.SectionCode
}

func UserModelToUserDTO(user *User, mUser *model.User) {
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
	user.DesiredRounds = make([]DesiredRound, len(mUser.DesiredRounds))
	user.InterestedFaculties = make([]FacultyInfo, len(mUser.InterestedFaculties))

	for i, round := range mUser.DesiredRounds {
		DesiredRoundModelToDTO(&user.DesiredRounds[i], &round)
	}

	for i, faculty := range mUser.InterestedFaculties {
		InterestedFacultyToFacultyInfo(&user.InterestedFaculties[i], &faculty)
	}
}

func InterestedFacultyToFacultyInfo(facultyInfo *FacultyInfo, m *model.InterestedFaculty) {
	facultyInfo.Department.Code = m.DepartmentCode
	facultyInfo.Department.Name.En = m.Department.Name.En
	facultyInfo.Department.Name.Th = m.Department.Name.Th
	facultyInfo.Faculty.Code = m.FacultyCode
	facultyInfo.Faculty.Name.En = m.Faculty.Name.En
	facultyInfo.Faculty.Name.Th = m.Faculty.Name.Th
	facultyInfo.Section.Code = m.SectionCode
	facultyInfo.Section.Name.En = m.Section.Name.En
	facultyInfo.Section.Name.Th = m.Section.Name.Th
}

func DesiredRoundModelToDTO(desiredRound *DesiredRound, m *model.DesiredRound) {
	desiredRound.Order = m.Order
	desiredRound.Round = string(m.Round)
}
