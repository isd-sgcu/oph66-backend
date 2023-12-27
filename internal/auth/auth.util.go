package auth

import "github.com/isd-sgcu/oph66-backend/internal/model"

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

	for i, desiredRound := range dto.DesiredRounds {
		ConvertDesiredInfoToDesiredRound(&desiredRound, user, &user.DesiredRounds[i])
	}

	for i, interestedFaculty := range dto.InterestedFaculties {
		ConvertFacultyInfoToInterestedFaculty(&interestedFaculty, user, &user.InterestedFaculties[i])
	}

	return user
}

func ConvertDesiredInfoToDesiredRound(dto *DesiredInfo, user *model.User, desiredRound *model.DesiredRound) {
	desiredRound.Order = dto.Order
	desiredRound.RoundCode = dto.Code
}

func ConvertFacultyInfoToInterestedFaculty(dto *FacultyInfo, user *model.User, interestedFaculty *model.InterestedFaculty) {
	interestedFaculty.Order = dto.Order
}
