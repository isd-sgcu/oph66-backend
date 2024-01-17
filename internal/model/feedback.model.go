package model

type FirstpartAnswer string

type SecondPartAnswer string

type Feedback struct {
	User    User             `gorm:"foreignKey:UserId;references:Id"`
	UserId  int              `gorm:""`
	Q1      FirstpartAnswer  `gorm:""`
	Q2      FirstpartAnswer  `gorm:""`
	Q3      FirstpartAnswer  `gorm:""`
	Q4      FirstpartAnswer  `gorm:""`
	Q5      FirstpartAnswer  `gorm:""`
	Q6      FirstpartAnswer  `gorm:""`
	Q7      SecondPartAnswer `gorm:""`
	Q8      SecondPartAnswer `gorm:""`
	Q9      SecondPartAnswer `gorm:""`
	Q10     SecondPartAnswer `gorm:""`
	Q11     SecondPartAnswer `gorm:""`
	Q12     SecondPartAnswer `gorm:""`
	Q13     SecondPartAnswer `gorm:""`
	Q14     SecondPartAnswer `gorm:""`
	Q15     SecondPartAnswer `gorm:""`
	Q16     SecondPartAnswer `gorm:""`
	Q17     SecondPartAnswer `gorm:""`
	Q18     SecondPartAnswer `gorm:""`
	Q19     SecondPartAnswer `gorm:""`
	Comment string           `gorm:""`
}
