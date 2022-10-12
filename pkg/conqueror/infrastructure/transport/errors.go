package transport

import (
	"net/http"

	"conqueror/pkg/conqueror/domain"
)

func mapErrorToStatus(err error) int {
	switch err {
	case domain.ErrLoginLength,
		domain.ErrNicknameLength:
		return http.StatusBadRequest
	case domain.ErrUserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
