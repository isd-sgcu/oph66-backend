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
	BadRequest            = &AppError{"bad-request", http.StatusBadRequest}
	InvalidEventId        = &AppError{"invalid-event-id", http.StatusNotFound}
	InvalidToken          = &AppError{"invalid-token", http.StatusUnauthorized}
	InvalidEmail          = &AppError{"invalid-email", http.StatusNotFound}
	DuplicateEmail        = &AppError{"duplicate-email", http.StatusConflict}
	Unauthorized          = &AppError{"unauthorized", http.StatusUnauthorized}
	ServiceUnavailable    = &AppError{"service-unavailable", http.StatusServiceUnavailable}
	UserNotFound          = &AppError{"user-not-found", http.StatusNotFound}
)
