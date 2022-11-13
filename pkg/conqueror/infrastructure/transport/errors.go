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
		domain.ErrTaskTitleLength,
		domain.ErrTaskDescriptionLength,
		domain.ErrDuplicateTaskTags,
		domain.ErrTaskTagNameLength,
		domain.ErrNoteTitleLength,
		domain.ErrNoteContentLength,
		domain.ErrNoteTagNameLength,
		domain.ErrDuplicateNoteTags:
		return http.StatusBadRequest
	case domain.ErrUserNotFound,
		domain.ErrSubjectNotFound,
		domain.ErrTaskNotFound,
		domain.ErrTaskTagNotFound,
		domain.ErrNoteNotFound,
		domain.ErrNoteTagNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
