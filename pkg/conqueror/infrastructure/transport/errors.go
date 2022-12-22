package transport

import (
	"net/http"

	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/service"

	"conqueror/pkg/conqueror/domain"
)

func mapErrorToStatus(err error) int {
	switch errors.Cause(err) {
	case uuid.ErrInvalidUUID,
		domain.ErrLoginLength,
		service.ErrUserAlreadyExists,
		service.ErrWeakPassword,
		domain.ErrSubjectTitleLength,
		domain.ErrTaskTitleLength,
		domain.ErrTaskDescriptionLength,
		domain.ErrDuplicateTaskTags,
		domain.ErrTaskTagNameLength,
		domain.ErrNoteTitleLength,
		domain.ErrNoteContentLength,
		domain.ErrNoteTagNameLength,
		domain.ErrDuplicateNoteTags,
		domain.ErrInvalidTaskStatus:
		return http.StatusBadRequest
	case domain.ErrUserNotFound,
		domain.ErrSubjectNotFound,
		domain.ErrTaskNotFound,
		domain.ErrTaskTagNotFound,
		domain.ErrNoteNotFound,
		domain.ErrNoteTagNotFound:
		return http.StatusNotFound
	case ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
