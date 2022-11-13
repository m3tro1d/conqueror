package transport

import (
	"context"
	"net/http"

	"conqueror/pkg/common/md5"
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
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
	ChangeTaskTags(ctx *gin.Context) error
	RemoveTask(ctx *gin.Context) error

	CreateTaskTag(ctx *gin.Context) error
	ChangeTaskTagName(ctx *gin.Context) error
	RemoveTaskTag(ctx *gin.Context) error

	CreateNote(ctx *gin.Context) error
	ChangeNoteTitle(ctx *gin.Context) error
	ChangeNoteContent(ctx *gin.Context) error
	RemoveNote(ctx *gin.Context) error

	CreateNoteTag(ctx *gin.Context) error
	ChangeNoteTagName(ctx *gin.Context) error
	RemoveNoteTag(ctx *gin.Context) error

	ListTasks(ctx *gin.Context) error
	ListNotes(ctx *gin.Context) error
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

func (api *publicAPI) ChangeTaskTags(ctx *gin.Context) error {
	taskID, err := uuid.FromString(ctx.Param("taskID"))
	if err != nil {
		return err
	}

	var request changeTaskTagsRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	tags, err := uuid.FromStrings(request.Tags)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskService().ChangeTaskTags(taskID, tags)
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

func (api *publicAPI) CreateTaskTag(ctx *gin.Context) error {
	var request createTaskTagRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	// TODO: user ID from auth
	err = api.dependencyContainer.TaskTagService().CreateTaskTag(uuid.UUID{}, request.Name)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (api *publicAPI) ChangeTaskTagName(ctx *gin.Context) error {
	taskTagID, err := uuid.FromString(ctx.Param("taskTagID"))
	if err != nil {
		return err
	}

	var request changeTaskTagNameRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskTagService().ChangeTaskTagName(taskTagID, request.NewName)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) RemoveTaskTag(ctx *gin.Context) error {
	taskTagID, err := uuid.FromString(ctx.Param("taskTagID"))
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskTagService().RemoveTaskTag(taskTagID)
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

func (api *publicAPI) CreateNoteTag(ctx *gin.Context) error {
	var request createNoteTagRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	// TODO: user ID from auth
	err = api.dependencyContainer.NoteTagService().CreateNoteTag(uuid.UUID{}, request.Name)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (api *publicAPI) ChangeNoteTagName(ctx *gin.Context) error {
	noteTagID, err := uuid.FromString(ctx.Param("noteTagID"))
	if err != nil {
		return err
	}

	var request changeNoteTagNameRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteTagService().ChangeNoteTagName(noteTagID, request.NewName)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) RemoveNoteTag(ctx *gin.Context) error {
	noteTagID, err := uuid.FromString(ctx.Param("noteTagID"))
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteTagService().RemoveNoteTag(noteTagID)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

func (api *publicAPI) ListTasks(ctx *gin.Context) error {
	// TODO: get actual user context
	userCtx := auth.NewUserContext(context.Background(), uuid.UUID{})

	tasks, err := api.dependencyContainer.TaskQueryService().ListTasks(userCtx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listTasksResponse{
		Tasks: queryTasksToApi(tasks),
	})
	return nil
}

func (api *publicAPI) ListNotes(ctx *gin.Context) error {
	// TODO: get actual user context
	userCtx := auth.NewUserContext(context.Background(), uuid.UUID{})

	notes, err := api.dependencyContainer.NoteQueryService().ListNotes(userCtx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listNotesResponse{
		Notes: queryNotesToApi(notes),
	})
	return nil
}
