package transport

import (
	"net/http"

	"conqueror/pkg/conqueror/infrastructure"

	"github.com/gin-gonic/gin"
)

type PublicAPI interface {
	RegisterUser(ctx *gin.Context) error
	CreateSubject(ctx *gin.Context) error
}

func NewPublicAPI(dependencyContainer infrastructure.DependencyContainer) PublicAPI {
	return &publicAPI{
		dependencyContainer: dependencyContainer,
	}
}

type publicAPI struct {
	dependencyContainer infrastructure.DependencyContainer
}

func (api *publicAPI) RegisterUser(ctx *gin.Context) error {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register user",
	})

	return nil
}

func (api *publicAPI) CreateSubject(ctx *gin.Context) error {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create subject",
	})

	return nil
}
