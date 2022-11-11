package transport

import (
	"context"
	"net/http"

	"conqueror/pkg/common/md5"
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/infrastructure"

	"github.com/gin-gonic/gin"
)

type PublicAPI interface {
	RegisterUser(ctx *gin.Context) error
	LoginUser(ctx *gin.Context) error

	CreateSubject(ctx *gin.Context) error
	ChangeSubjectTitle(ctx *gin.Context) error
	RemoveSubject(ctx *gin.Context) error

	CreateTask(ctx *gin.Context) error
	ChangeTaskTitle(ctx *gin.Context) error
	ChangeTaskDescription(ctx *gin.Context) error
	RemoveTask(ctx *gin.Context) error

	CreateNote(ctx *gin.Context) error
	ChangeNoteTitle(ctx *gin.Context) error
	ChangeNoteContent(ctx *gin.Context) error
	RemoveNote(ctx *gin.Context) error
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
	var request registerUserRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.UserService().RegisterUser(request.Login, request.Password, request.Nickname)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (api *publicAPI) LoginUser(ctx *gin.Context) error {
	var request loginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	user, err := api.dependencyContainer.UserQueryService().GetByLogin(context.Background(), request.Login)
	if err != nil {
		return err
	}

	if md5.Hash(request.Password) != user.Password {
		ctx.Status(http.StatusForbidden)
		return nil
	}

	// TODO: generate and send token

	return nil
}

func (api *publicAPI) CreateSubject(ctx *gin.Context) error {
	var request createSubjectRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	// TODO: user ID from auth
	err = api.dependencyContainer.SubjectService().CreateSubject(uuid.UUID{}, request.Title)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (api *publicAPI) ChangeSubjectTitle(ctx *gin.Context) error {
	subjectID, err := uuid.FromString(ctx.Param("subjectID"))
	if err != nil {
		return err
	}

	var request changeSubjectTitleRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.SubjectService().ChangeSubjectTitle(subjectID, request.NewTitle)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) RemoveSubject(ctx *gin.Context) error {
	subjectID, err := uuid.FromString(ctx.Param("subjectID"))
	if err != nil {
		return err
	}

	err = api.dependencyContainer.SubjectService().RemoveSubject(subjectID)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

func (api *publicAPI) CreateTask(ctx *gin.Context) error {
	var request createTaskRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	subjectID, err := uuid.OptionalFromString(request.SubjectID)
	if err != nil {
		return err
	}

	// TODO: user ID from auth
	err = api.dependencyContainer.TaskService().CreateTask(
		uuid.UUID{},
		request.DueDate,
		request.Title,
		request.Description,
		subjectID,
	)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (api *publicAPI) ChangeTaskTitle(ctx *gin.Context) error {
	taskID, err := uuid.FromString(ctx.Param("taskID"))
	if err != nil {
		return err
	}

	var request changeTaskTitleRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskService().ChangeTaskTitle(taskID, request.NewTitle)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) ChangeTaskDescription(ctx *gin.Context) error {
	taskID, err := uuid.FromString(ctx.Param("taskID"))
	if err != nil {
		return err
	}

	var request changeTaskDescriptionRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskService().ChangeTaskDescription(taskID, request.NewDescription)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) RemoveTask(ctx *gin.Context) error {
	taskID, err := uuid.FromString(ctx.Param("taskID"))
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskService().RemoveTask(taskID)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

func (api *publicAPI) CreateNote(ctx *gin.Context) error {
	var request createNoteRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	subjectID, err := uuid.OptionalFromString(request.SubjectID)
	if err != nil {
		return err
	}

	// TODO: user ID from auth
	err = api.dependencyContainer.NoteService().CreateNote(
		uuid.UUID{},
		request.Title,
		request.Content,
		subjectID,
	)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (api *publicAPI) ChangeNoteTitle(ctx *gin.Context) error {
	noteID, err := uuid.FromString(ctx.Param("noteID"))
	if err != nil {
		return err
	}

	var request changeNoteTitleRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteService().ChangeNoteTitle(noteID, request.NewTitle)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) ChangeNoteContent(ctx *gin.Context) error {
	noteID, err := uuid.FromString(ctx.Param("noteID"))
	if err != nil {
		return err
	}

	var request changeNoteContentRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteService().ChangeNoteContent(noteID, request.NewContent)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) RemoveNote(ctx *gin.Context) error {
	noteID, err := uuid.FromString(ctx.Param("noteID"))
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteService().RemoveNote(noteID)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}
