package apperror

import (
	"net/http"
)

type AppError struct {
	Id       string
	HttpCode int
}

func (a *AppError) Error() string {
	return a.Id
}

var (
	InvalidFeatureFlagKey = &AppError{"invalid-feature-flag-key", http.StatusNotFound}
	InternalError         = &AppError{"internal-server-error", http.StatusInternalServerError}
	InvalidEventId        = &AppError{"invalid-event-id", http.StatusNotFound}
)
