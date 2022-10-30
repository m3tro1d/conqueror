package transport

import (
	"context"
	stderrors "errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
)

const authHeader = "X-Auth-Token"

func (api *publicAPI) getUserContext(ctx *gin.Context) (auth.UserContext, error) {
	tokenString := ctx.GetHeader(authHeader)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, stderrors.New("invalid signing method")
		}

		return []byte("boobz"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := uuid.FromString(fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {

			return nil, err
		}

		return auth.NewUserContext(context.Background(), userId), nil
	}

	return nil, stderrors.New("invalid token")
}
