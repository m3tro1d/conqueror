package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/app"
	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/infrastructure/mysql"

	"github.com/jmoiron/sqlx"
)

type DependencyContainer interface {
	UserService() app.UserService
	SubjectService() app.SubjectService

	UserQueryService() query.UserQueryService
}

func NewDependencyContainer(ctx context.Context, db *sqlx.DB) (DependencyContainer, error) {
	conn, err := db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	userRepository := mysql.NewUserRepository(ctx, conn)
	userService := app.NewUserService(userRepository)

	subjectRepository := mysql.NewSubjectRepository(ctx, conn)
	subjectService := app.NewSubjectService(subjectRepository, userRepository)

	userQueryService := mysql.NewUserQueryService(ctx, conn)

	return &dependencyContainer{
		userService:    userService,
		subjectService: subjectService,

		userQueryService: userQueryService,
	}, nil
}

type dependencyContainer struct {
	userService    app.UserService
	subjectService app.SubjectService

	userQueryService query.UserQueryService
}

func (c *dependencyContainer) UserService() app.UserService {
	return c.userService
}

func (c *dependencyContainer) SubjectService() app.SubjectService {
	return c.subjectService
}

func (c *dependencyContainer) UserQueryService() query.UserQueryService {
	return c.userQueryService
}
