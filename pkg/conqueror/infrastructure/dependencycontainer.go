package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/app"
	"conqueror/pkg/conqueror/infrastructure/mysql"

	"github.com/jmoiron/sqlx"
)

type DependencyContainer interface {
	UserService() app.UserService
}

func NewDependencyContainer(ctx context.Context, db *sqlx.DB) (DependencyContainer, error) {
	conn, err := db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	userRepository := mysql.NewUserRepository(ctx, conn)
	userService := app.NewUserService(userRepository)

	return &dependencyContainer{
		userService: userService,
	}, nil
}

type dependencyContainer struct {
	userService app.UserService
}

func (c *dependencyContainer) UserService() app.UserService {
	return c.userService
}
