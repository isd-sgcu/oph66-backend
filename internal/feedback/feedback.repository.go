package feedback

import (
	"github.com/isd-sgcu/oph66-backend/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateFeedback(feedback *model.Feedback) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}
func (r *repositoryImpl) CreateFeedback(feedback *model.Feedback) error {
	return r.db.Model(feedback).Create(feedback).Error
}
