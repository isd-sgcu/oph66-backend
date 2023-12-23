package featureflag

import "github.com/isd-sgcu/oph66-backend/database"

type FeatureFlag struct {
	Key           string         `json:"key" gorm:"primaryKey"`
	Enabled       bool           `json:"enabled"`
	CacheDuration int            `json:"-"`
	ExtraInfo     database.JSONB `json:"extra_info"`
}

func (m FeatureFlag) TableName() string {
	return "feature_flags"
}
