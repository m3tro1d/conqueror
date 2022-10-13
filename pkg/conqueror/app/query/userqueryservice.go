package query

import "conqueror/pkg/common/uuid"

type UserQueryService interface {
	GetByLogin(login string) (User, error)
}

type User struct {
	UserID   uuid.UUID
	Login    string
	Password string
	Nickname string
}
