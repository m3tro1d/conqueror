package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(publicAPI PublicAPI) http.Handler {
	router := gin.Default()

	router.POST("/api/v1/user", handlerFunc(publicAPI.RegisterUser))
	router.POST("/api/v1/user/login", handlerFunc(publicAPI.LoginUser))

	router.POST("/api/v1/subject", handlerFunc(publicAPI.CreateSubject))
	router.PUT("/api/v1/subject/:subjectID", handlerFunc(publicAPI.ChangeSubjectTitle))
	router.DELETE("/api/v1/subject/:subjectID", handlerFunc(publicAPI.RemoveSubject))

	return router
}

type handler = func(ctx *gin.Context) error

func handlerFunc(handler handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := handler(ctx)
		processError(ctx, err)
	}
}

func processError(ctx *gin.Context, err error) {
	if err != nil {
		status := mapErrorToStatus(err)
		if status == http.StatusInternalServerError {
			ctx.String(status, "%+v", err)
			return
		}

		ctx.String(status, err.Error())
	}
}
