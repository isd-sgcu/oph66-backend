package featureflag

type response struct {
	Key       string `json:"key" example:"livestream"`
	Enabled   bool   `json:"enabled" example:"true"`
	ExtraInfo string `json:"extra_info" example:"https://www.youtube.com/watch?v=6n3pFFPSlW4"`
}

type errorResponse struct {
	Instance string `json:"instance" example:"/featureflag/live"`	
	Title    string `json:"title" example:"internal-server-error"`
}

type invalidResponse struct {
	Instance string `json:"instance" example:"/featureflag/live"`
	Title    string `json:"title" example:"invalid-feature-flag-key"`
}