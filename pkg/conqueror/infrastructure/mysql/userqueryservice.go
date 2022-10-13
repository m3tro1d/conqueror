package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/domain"
)

func NewUserQueryService(ctx context.Context, client *sqlx.Conn) query.UserQueryService {
	return &userQueryService{
		ctx:    ctx,
		client: client,
	}
}

type userQueryService struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (s *userQueryService) GetByLogin(login string) (query.User, error) {
	const sqlQuery = `SELECT id, login, password, nickname
		              FROM user
		              WHERE login = ?
		              LIMIT 1`

	var user sqlxUser
	err := s.client.GetContext(s.ctx, &user, sqlQuery, login)
	if err == sql.ErrNoRows {
		return query.User{}, errors.WithStack(domain.ErrUserNotFound)
	} else if err != nil {
		return query.User{}, err
	}

	return query.User{
		UserID:   uuid.UUID(user.ID),
		Login:    user.Login,
		Password: user.Password,
		Nickname: user.Nickname,
	}, nil
}
