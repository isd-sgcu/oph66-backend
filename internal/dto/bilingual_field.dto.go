package dto

import "github.com/isd-sgcu/oph66-backend/internal/model"

type BilingualField struct {
	Th string `json:"th"`
	En string `json:"en"`
}

func BilingualModelToDTO(m *model.Bilingual) BilingualField {
	var bf BilingualField
	bf.En = m.En
	bf.Th = m.Th
	return bf
}
