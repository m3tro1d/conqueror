package query

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
)

type UserQueryService interface {
	GetCurrentUser(ctx auth.UserContext) (UserData, error)
	GetByLogin(login string) (UserData, error)
}

type UserData struct {
	UserID   uuid.UUID
	Login    string
	Password string
}
