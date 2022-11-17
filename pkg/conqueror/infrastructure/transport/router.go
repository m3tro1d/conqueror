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
	router.PUT("/api/v1/subject/:subjectID/title", handlerFunc(publicAPI.ChangeSubjectTitle))
	router.DELETE("/api/v1/subject/:subjectID", handlerFunc(publicAPI.RemoveSubject))

	router.POST("/api/v1/task", handlerFunc(publicAPI.CreateTask))
	router.PUT("/api/v1/task/:taskID/title", handlerFunc(publicAPI.ChangeTaskTitle))
	router.PUT("/api/v1/task/:taskID/tags", handlerFunc(publicAPI.ChangeTaskTags))
	router.PUT("/api/v1/task/:taskID/description", handlerFunc(publicAPI.ChangeTaskDescription))
	router.DELETE("/api/v1/task/:taskID", handlerFunc(publicAPI.RemoveTask))

	router.POST("/api/v1/tasktag", handlerFunc(publicAPI.CreateTaskTag))
	router.PUT("/api/v1/tasktag/:taskTagID/name", handlerFunc(publicAPI.ChangeTaskTagName))
	router.DELETE("/api/v1/tasktag/:taskTagID", handlerFunc(publicAPI.RemoveTaskTag))

	router.POST("/api/v1/note", handlerFunc(publicAPI.CreateNote))
	router.PUT("/api/v1/note/:noteID/title", handlerFunc(publicAPI.ChangeNoteTitle))
	router.PUT("/api/v1/note/:noteID/content", handlerFunc(publicAPI.ChangeNoteContent))
	router.PUT("/api/v1/note/:noteID/tags", handlerFunc(publicAPI.ChangeNoteTags))
	router.DELETE("/api/v1/note/:noteID", handlerFunc(publicAPI.RemoveNote))

	router.POST("/api/v1/notetag", handlerFunc(publicAPI.CreateNoteTag))
	router.PUT("/api/v1/notetag/:noteTagID/name", handlerFunc(publicAPI.ChangeNoteTagName))
	router.DELETE("/api/v1/notetag/:noteTagID", handlerFunc(publicAPI.RemoveNoteTag))

	router.GET("/api/v1/tasks", handlerFunc(publicAPI.ListTasks))
	router.GET("/api/v1/task/tags", handlerFunc(publicAPI.ListTaskTags))
	router.GET("/api/v1/notes", handlerFunc(publicAPI.ListNotes))
	router.GET("/api/v1/note/tags", handlerFunc(publicAPI.ListNoteTags))

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
