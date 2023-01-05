package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(publicAPI PublicAPI) http.Handler {
	router := gin.Default()

	router.POST("/api/v1/user", handlerFunc(publicAPI.RegisterUser))
	router.POST("/api/v1/user/login", handlerFunc(publicAPI.LoginUser))
	router.PATCH("/api/v1/user/avatar", handlerFunc(publicAPI.ChangeUserAvatar))
	router.GET("/api/v1/user", handlerFunc(publicAPI.GetUser))

	router.POST("/api/v1/subject", handlerFunc(publicAPI.CreateSubject))
	router.PATCH("/api/v1/subject/:subjectID/title", handlerFunc(publicAPI.ChangeSubjectTitle))
	router.DELETE("/api/v1/subject/:subjectID", handlerFunc(publicAPI.RemoveSubject))

	router.POST("/api/v1/task", handlerFunc(publicAPI.CreateTask))
	router.PATCH("/api/v1/task/:taskID", handlerFunc(publicAPI.UpdateTask))
	router.PATCH("/api/v1/task/:taskID/status", handlerFunc(publicAPI.ChangeTaskStatus))
	router.DELETE("/api/v1/task/:taskID", handlerFunc(publicAPI.RemoveTask))

	router.POST("/api/v1/note", handlerFunc(publicAPI.CreateNote))
	router.PATCH("/api/v1/note/:noteID", handlerFunc(publicAPI.UpdateNote))
	router.DELETE("/api/v1/note/:noteID", handlerFunc(publicAPI.RemoveNote))

	router.GET("/api/v1/subjects", handlerFunc(publicAPI.ListSubjects))
	router.GET("/api/v1/tasks", handlerFunc(publicAPI.ListTasks))
	router.GET("/api/v1/task/:taskID", handlerFunc(publicAPI.GetTask))
	router.GET("/api/v1/notes", handlerFunc(publicAPI.ListNotes))
	router.GET("/api/v1/note/:noteID", handlerFunc(publicAPI.GetNote))

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
