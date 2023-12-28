package dto

type response struct {
	Key       string `example:"livestream"                                  json:"key"`
	Enabled   bool   `example:"true"                                        json:"enabled"`
	ExtraInfo string `example:"https://www.youtube.com/watch?v=6n3pFFPSlW4" json:"extra_info"`
}

type errorResponse struct {
	Instance string `example:"/featureflag/live"     json:"instance"`
	Title    string `example:"internal-server-error" json:"title"`
}

type invalidResponse struct {
	Instance string `example:"/featureflag/live"        json:"instance"`
	Title    string `example:"invalid-feature-flag-key" json:"title"`
}
