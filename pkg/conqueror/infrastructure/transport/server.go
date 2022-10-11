package transport

import (
	"net/http"

	"conqueror/pkg/conqueror/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewPublicAPI(dependencyContainer infrastructure.DependencyContainer) *PublicAPI {
	return &PublicAPI{
		dependencyContainer: dependencyContainer,
	}
}

type PublicAPI struct {
	dependencyContainer infrastructure.DependencyContainer
}

func (api *PublicAPI) RegisterUser(ctx *gin.Context) error {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hi",
	})

	return nil
}
