package featureflag

import "gorm.io/gorm"

type Repository interface {
	FindOneByKey(result *FeatureFlag, key string) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

var _ Repository = &repositoryImpl{}

type repositoryImpl struct {
	db *gorm.DB
}

func (r *repositoryImpl) FindOneByKey(result *FeatureFlag, key string) error {
	return r.db.Model(result).First(result, "key = ?", key).Error
}
