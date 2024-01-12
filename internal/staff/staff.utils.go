package staff

import (
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func UserModelToCheckinUser(m *model.User) dto.AttendeeStaffCheckinUser {
	var checkinUser dto.AttendeeStaffCheckinUser
	if m == nil {
		return checkinUser
	}
	checkinUser.FirstName = m.FirstName
	checkinUser.LastName = m.LastName
	checkinUser.Allergies = m.Allergies
	checkinUser.MedicalCondition = m.MedicalCondition

	return checkinUser
}
