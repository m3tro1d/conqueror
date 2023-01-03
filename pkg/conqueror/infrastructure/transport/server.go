package transport

import (
	"context"
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"conqueror/pkg/common/md5"
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/infrastructure"

	"github.com/gin-gonic/gin"
)

const (
	authHeader  = "X-Auth-Token"
	userIDClaim = "user_id"
)

var (
	ErrUnauthorized = stderrors.New("unauthorized")
)

type PublicAPI interface {
	RegisterUser(ctx *gin.Context) error
	LoginUser(ctx *gin.Context) error
	ChangeUserAvatar(ctx *gin.Context) error
	GetUser(ctx *gin.Context) error

	CreateSubject(ctx *gin.Context) error
	ChangeSubjectTitle(ctx *gin.Context) error
	RemoveSubject(ctx *gin.Context) error

	CreateTask(ctx *gin.Context) error
	ChangeTaskTitle(ctx *gin.Context) error
	ChangeTaskDescription(ctx *gin.Context) error
	ChangeTaskStatus(ctx *gin.Context) error
	ChangeTaskTags(ctx *gin.Context) error
	RemoveTask(ctx *gin.Context) error

	CreateTaskTag(ctx *gin.Context) error
	ChangeTaskTagName(ctx *gin.Context) error
	RemoveTaskTag(ctx *gin.Context) error

	CreateNote(ctx *gin.Context) error
	ChangeNoteTitle(ctx *gin.Context) error
	ChangeNoteContent(ctx *gin.Context) error
	ChangeNoteTags(ctx *gin.Context) error
	RemoveNote(ctx *gin.Context) error

	CreateNoteTag(ctx *gin.Context) error
	ChangeNoteTagName(ctx *gin.Context) error
	RemoveNoteTag(ctx *gin.Context) error

	ListSubjects(ctx *gin.Context) error
	ListTasks(ctx *gin.Context) error
	ListTaskTags(ctx *gin.Context) error
	ListNotes(ctx *gin.Context) error
	ListNoteTags(ctx *gin.Context) error
}

func NewPublicAPI(dependencyContainer infrastructure.DependencyContainer, secret []byte) PublicAPI {
	return &publicAPI{
		dependencyContainer: dependencyContainer,
		secret:              secret,
	}
}

type publicAPI struct {
	dependencyContainer infrastructure.DependencyContainer
	secret              []byte
}

func (api *publicAPI) RegisterUser(ctx *gin.Context) error {
	var request registerUserRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.UserService().RegisterUser(request.Login, request.Password)
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

	user, err := api.dependencyContainer.UserQueryService().GetByLogin(request.Login)
	if err != nil {
		return err
	}

	if md5.Hash(request.Password) != user.Password {
		ctx.Status(http.StatusForbidden)
		return nil
	}

	token, err := api.generateToken(user.UserID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, loginResponse{
		Token: token,
	})
	return nil
}

func (api *publicAPI) ChangeUserAvatar(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	requestFile, err := ctx.FormFile("avatar")
	if err != nil {
		return err
	}

	file, err := requestFile.Open()
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	err = api.dependencyContainer.UserService().ChangeUserAvatar(userCtx.UserID(), file)
	if err != nil {
		return err
	}

	return nil
}

func (api *publicAPI) GetUser(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	user, err := api.dependencyContainer.UserQueryService().GetCurrentUser(userCtx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, queryUserToApi(user))
	return nil
}

func (api *publicAPI) CreateSubject(ctx *gin.Context) error {
	var request createSubjectRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.SubjectService().CreateSubject(userCtx.UserID(), request.Title)
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

	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskService().CreateTask(
		userCtx.UserID(),
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

func (api *publicAPI) ChangeTaskStatus(ctx *gin.Context) error {
	taskID, err := uuid.FromString(ctx.Param("taskID"))
	if err != nil {
		return err
	}

	var request changeTaskStatusRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskService().ChangeTaskStatus(taskID, request.NewStatus)
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

	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.TaskTagService().CreateTaskTag(userCtx.UserID(), request.Name)
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

	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteService().CreateNote(
		userCtx.UserID(),
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

func (api *publicAPI) ChangeNoteTags(ctx *gin.Context) error {
	noteID, err := uuid.FromString(ctx.Param("noteID"))
	if err != nil {
		return err
	}

	var request changeNoteTagsRequest
	err = ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	tags, err := uuid.FromStrings(request.Tags)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteService().ChangeNoteTags(noteID, tags)
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

	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	err = api.dependencyContainer.NoteTagService().CreateNoteTag(userCtx.UserID(), request.Name)
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

func (api *publicAPI) ListSubjects(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	subjects, err := api.dependencyContainer.SubjectQueryService().ListSubjects(userCtx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listSubjectsResponse{
		Subjects: querySubjectsToApi(subjects),
	})
	return nil
}

func (api *publicAPI) ListTasks(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	spec := buildListTasksSpecification(ctx)
	tasks, err := api.dependencyContainer.TaskQueryService().ListTasks(userCtx, spec)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listTasksResponse{
		Tasks: queryTasksToApi(tasks),
	})
	return nil
}

func (api *publicAPI) ListTaskTags(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	tags, err := api.dependencyContainer.TaskQueryService().ListTaskTags(userCtx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listTaskTagsResponse{
		Tags: queryTaskTagsToApi(tags),
	})
	return nil
}

func (api *publicAPI) ListNotes(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	spec := buildListNotesSpecification(ctx)
	notes, err := api.dependencyContainer.NoteQueryService().ListNotes(userCtx, spec)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listNotesResponse{
		Notes: queryNotesToApi(notes),
	})
	return nil
}

func (api *publicAPI) ListNoteTags(ctx *gin.Context) error {
	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	tags, err := api.dependencyContainer.NoteQueryService().ListNoteTags(userCtx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, listNoteTagsResponse{
		Tags: queryNoteTagsToApi(tags),
	})
	return nil
}

func (api *publicAPI) generateToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDClaim: userID.String(),
	})

	return token.SignedString(api.secret)
}

func (api *publicAPI) getUserContext(ctx *gin.Context) (auth.UserContext, error) {
	tokenString := ctx.GetHeader(authHeader)
	if len(tokenString) == 0 {
		ctx.Status(http.StatusUnauthorized)
		return nil, errors.WithStack(ErrUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return api.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, stderrors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, stderrors.New("invalid token claims")
	}

	userIDString := claims[userIDClaim].(string)
	userID, err := uuid.FromString(userIDString)
	if err != nil {
		return nil, err
	}

	return auth.NewUserContext(context.Background(), userID), nil
}
