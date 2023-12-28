package dto

import "encoding/json"

type JSON any

type FeatureFlagResponse struct {
	Key           string          `example:"livestream"                                  json:"key"`
	Enabled       bool            `example:"true"                                        json:"enabled"`
	ExtraInfo     json.RawMessage `example:"<jsonobject>" swaggertype:"string" json:"extra_info"`
	CacheDuration int             `json:"-"`
}

type FeatureFlagInternalErrorResponse struct {
	Instance string `example:"/featureflag/live"     json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type FeatureFlagInvalidKeyResponse struct {
	Instance string `example:"/featureflag/live"        json:"instance"`
	Title    string `example:"invalid-feature-flag-key" json:"title"`
}
