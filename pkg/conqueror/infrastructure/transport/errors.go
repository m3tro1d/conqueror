package transport

import (
	"net/http"

	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"

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
		domain.ErrSubjectTitleLength,
		domain.ErrNoteTitleLength,
		domain.ErrNoteContentLength:
		return http.StatusBadRequest
	case domain.ErrUserNotFound,
		domain.ErrSubjectNotFound,
		domain.ErrTaskNotFound,
		domain.ErrNoteNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
