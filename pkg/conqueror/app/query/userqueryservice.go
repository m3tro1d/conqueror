package query

import (
	"context"

	"conqueror/pkg/common/uuid"
)

type UserQueryService interface {
	GetByLogin(ctx context.Context, login string) (UserData, error)
}

type UserData struct {
	UserID   uuid.UUID
	Login    string
	Password string
	Nickname string
}
