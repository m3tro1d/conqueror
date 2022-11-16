package transport

import (
	"context"
	stderrors "errors"

	"github.com/gin-gonic/gin"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
)

const authHeader = "X-Auth-ID"

func (api *publicAPI) getUserContext(ctx *gin.Context) (auth.UserContext, error) {
	userIDString := ctx.GetHeader(authHeader)

	if len(userIDString) == 0 {
		return nil, stderrors.New("invalid user ID")
	}

	userID, err := uuid.FromString(userIDString)
	if err != nil {
		return nil, err
	}

	return auth.NewUserContext(context.Background(), userID), nil
}
