package transport

import (
	"context"
	stderrors "errors"
	"fmt"
	"net/http"

	"conqueror/pkg/common/md5"
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/infrastructure"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type PublicAPI interface {
	RegisterUser(ctx *gin.Context) error
	LoginUser(ctx *gin.Context) error

	CreateSubject(ctx *gin.Context) error
	ChangeSubjectTitle(ctx *gin.Context) error
	RemoveSubject(ctx *gin.Context) error
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

	userCtx, err := api.getUserContext(ctx)
	if err != nil {
		return err
	}

	user, err := api.dependencyContainer.UserQueryService().GetByLogin(userCtx, request.Login)
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

func (api *publicAPI) getUserContext(ctx *gin.Context) (auth.UserContext, error) {
	tokenString := ctx.GetHeader("X-Auth-Token")

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

	// TODO: add an error and handle in errors.go
	return nil, stderrors.New("invalid token")
}
