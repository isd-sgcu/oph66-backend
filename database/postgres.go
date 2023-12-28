package database

import (
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(config *cfgldr.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(config.DatabaseConfig.Url), &gorm.Config{TranslateError: true})
}
