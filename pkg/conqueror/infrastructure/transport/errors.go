package transport

import (
	"net/http"

	"conqueror/pkg/common/uuid"
	"github.com/pkg/errors"

	"conqueror/pkg/conqueror/app"
	"conqueror/pkg/conqueror/domain"
)

func mapErrorToStatus(err error) int {
	switch errors.Cause(err) {
	case uuid.ErrInvalidUUID,
		domain.ErrLoginLength,
		domain.ErrNicknameLength,
		app.ErrUserAlreadyExists,
		app.ErrWeakPassword,
		domain.ErrSubjectTitleLength:
		return http.StatusBadRequest
	case domain.ErrUserNotFound,
		domain.ErrSubjectNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
