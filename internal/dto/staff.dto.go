package dto

type AttendeeStaffCheckinResponse struct {
	User           *AttendeeStaffCheckinUser `json:"user"`
	AlreadyCheckin bool                      `json:"already_checkin"`
}

type AttendeeStaffCheckinUser struct {
	FirstName        string `json:"first_name"        example:"John"`
	LastName         string `json:"last_name"         example:"Doe"`
	Allergies        string `json:"allergies"         example:"Romantic"`
	MedicalCondition string `json:"medical_condition" example:"Unlovable"`
}
