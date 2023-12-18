package featureflag

type FeatureFlag struct {
	Key           string `gorm:"primaryKey"`
	Value         bool
	CacheDuration int
}

func (m FeatureFlag) TableName() string {
	return "feature_flags"
}
