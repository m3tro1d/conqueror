package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/app"
	"conqueror/pkg/conqueror/infrastructure/mysql"

	"github.com/jmoiron/sqlx"
)

func NewDependencyContainer(ctx context.Context, db *sqlx.DB) (*DependencyContainer, error) {
	conn, err := db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	userRepository := mysql.NewUserRepository(ctx, conn)
	userService := app.NewUserService(userRepository)

	return &DependencyContainer{
		userService: userService,
	}, nil
}

type DependencyContainer struct {
	userService app.UserService
}

func (c *DependencyContainer) UserService() app.UserService {
	return c.userService
}
