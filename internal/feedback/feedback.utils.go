package feedback

import (
	"github.com/isd-sgcu/oph66-backend/internal/dto"
	"github.com/isd-sgcu/oph66-backend/internal/model"
)

func FeedbackDTOToModel(dto *dto.SubmitFeedbackDTO) model.Feedback {
	var m model.Feedback
	if dto == nil {
		return m
	}
	m.Q1 = model.FirstpartAnswer(dto.Q1)
	m.Q2 = model.FirstpartAnswer(dto.Q2)
	m.Q3 = model.FirstpartAnswer(dto.Q3)
	m.Q4 = model.FirstpartAnswer(dto.Q4)
	m.Q5 = model.FirstpartAnswer(dto.Q5)
	m.Q6 = model.FirstpartAnswer(dto.Q6)
	m.Q7 = model.SecondPartAnswer(dto.Q7)
	m.Q8 = model.SecondPartAnswer(dto.Q8)
	m.Q9 = model.SecondPartAnswer(dto.Q9)
	m.Q10 = model.SecondPartAnswer(dto.Q10)
	m.Q11 = model.SecondPartAnswer(dto.Q11)
	m.Q12 = model.SecondPartAnswer(dto.Q12)
	m.Q13 = model.SecondPartAnswer(dto.Q13)
	m.Q14 = model.SecondPartAnswer(dto.Q14)
	m.Q15 = model.SecondPartAnswer(dto.Q15)
	m.Q16 = model.SecondPartAnswer(dto.Q16)
	m.Q17 = model.SecondPartAnswer(dto.Q17)
	m.Q18 = model.SecondPartAnswer(dto.Q18)
	m.Q19 = model.SecondPartAnswer(dto.Q19)
	m.Comment = dto.Comment

	return m
}
