package auth

func ConvertRegisterRequestDTOToUser(dto *RegisterRequestDTO, email string) (user *User) {
	user = &User{}
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
	user.DesiredRounds = make([]DesiredRound, len(dto.DesiredRounds))
	user.InterestedFaculties = make([]InterestedFaculty, len(dto.InterestedFaculties))

	for i, desiredRound := range dto.DesiredRounds {
		ConvertDesiredInfoToDesiredRound(&desiredRound, user, &user.DesiredRounds[i])
	}

	for i, interestedFaculty := range dto.InterestedFaculties {
		ConvertFacultyInfoToInterestedFaculty(&interestedFaculty, user, &user.InterestedFaculties[i])
	}

	return user
}

func ConvertDesiredInfoToDesiredRound(dto *DesiredInfo, user *User, desiredRound *DesiredRound) {
	desiredRound.Order = dto.Order
	desiredRound.RoundCode = dto.Code
}

func ConvertFacultyInfoToInterestedFaculty(dto *FacultyInfo, user *User, interestedFaculty *InterestedFaculty) {
	interestedFaculty.Order = dto.Order
	interestedFaculty.FacultyCode = dto.FacultyCode
	interestedFaculty.DepartmentCode = dto.DepartmentCode
	interestedFaculty.SectionCode = dto.SectionCode
}