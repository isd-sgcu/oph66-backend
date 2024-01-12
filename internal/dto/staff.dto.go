package dto

type AttendeeStaffCheckinResponse struct {
	User           *AttendeeStaffCheckinUser `json:"user"`
	AlreadyCheckin bool                      `json:"already_checkin"`
}

type AttendeeStaffCheckinUser struct {
	FirstName        string `example:"John"      json:"first_name"`
	LastName         string `example:"Doe"       json:"last_name"`
	Allergies        string `example:"Romantic"  json:"allergies"`
	MedicalCondition string `example:"Unlovable" json:"medical_condition"`
}
