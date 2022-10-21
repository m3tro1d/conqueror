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

func NewUserQueryService(client *sqlx.Conn) query.UserQueryService {
	return &userQueryService{
		client: client,
	}
}

type userQueryService struct {
	client *sqlx.Conn
}

func (s *userQueryService) GetByLogin(ctx context.Context, login string) (query.User, error) {
	const sqlQuery = `SELECT id, login, password, nickname
		              FROM user
		              WHERE login = ?
		              LIMIT 1`

	var user sqlxUser
	err := s.client.GetContext(ctx, &user, sqlQuery, login)
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
