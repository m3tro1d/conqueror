package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/domain"
	"conqueror/pkg/conqueror/infrastructure/mysql"

	"github.com/jmoiron/sqlx"
)

func NewDependencyContainer(ctx context.Context, db *sqlx.DB) *DependencyContainer {
	return &DependencyContainer{
		ctx: ctx,
		db:  db,
	}
}

type DependencyContainer struct {
	ctx context.Context
	db  *sqlx.DB
}

func (c *DependencyContainer) userRepository() (domain.UserRepository, error) {
	conn, err := c.db.Connx(c.ctx)
	if err != nil {
		return nil, err
	}

	return mysql.NewUserRepository(c.ctx, conn), nil
}
