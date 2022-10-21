package auth

import (
	"context"

	"conqueror/pkg/common/uuid"
)

type UserContext interface {
	context.Context

	UserID() uuid.UUID
}

func NewUserContext(ctx context.Context, userID uuid.UUID) UserContext {
	return &userContext{
		Context: ctx,

		userID: userID,
	}
}

type userContext struct {
	context.Context

	userID uuid.UUID
}

func (ctx *userContext) UserID() uuid.UUID {
	return ctx.userID
}
