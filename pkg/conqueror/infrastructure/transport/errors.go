package transport

import (
	"net/http"

	"github.com/pkg/errors"

	"conqueror/pkg/conqueror/domain"
)

func mapErrorToStatus(err error) int {
	switch errors.Cause(err) {
	case domain.ErrLoginLength,
		domain.ErrNicknameLength:
		return http.StatusBadRequest
	case domain.ErrUserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
