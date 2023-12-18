package featureflag

import "gorm.io/gorm"

type Repository interface {
	FindOneByKey(result *FeatureFlag, key string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}
func (r *repositoryImpl) FindOneByKey(result *FeatureFlag, key string) error {
	return r.db.Model(result).First(result, "key = ?", key).Error
}
