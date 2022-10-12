package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(publicAPI PublicAPI) http.Handler {
	router := gin.Default()

	router.POST("/api/user", func(ctx *gin.Context) {
		err := publicAPI.RegisterUser(ctx)
		if err != nil {
			ctx.String(mapErrorToStatus(err), err.Error())
		}
	})

	return router
}
