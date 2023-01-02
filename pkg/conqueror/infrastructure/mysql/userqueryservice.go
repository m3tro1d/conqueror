package mysql

import (
	"conqueror/pkg/conqueror/app/auth"
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/domain"
)

const fileURLTemplate = "/files/%s"

func NewUserQueryService(client ClientContext, filesDir string) query.UserQueryService {
	return &userQueryService{
		client:   client,
		filesDir: filesDir,
	}
}

type userQueryService struct {
	client   ClientContext
	filesDir string
}

func (s *userQueryService) GetCurrentUser(ctx auth.UserContext) (query.UserData, error) {
	const sqlQuery = `SELECT u.id, u.login, u.password, i.id AS avatar_id, i.path AS avatar_path
		              FROM user u
		              	LEFT JOIN image i on u.avatar_id = i.id
		              WHERE u.id = ?
		              LIMIT 1`

	var user sqlxQueryUser
	err := s.client.GetContext(ctx, &user, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return query.UserData{}, errors.WithStack(domain.ErrUserNotFound)
	} else if err != nil {
		return query.UserData{}, err
	}

	var avatar *query.ImageData
	if user.AvatarID.ToOptionalUUID() != nil {
		avatar = &query.ImageData{
			ImageID: *user.AvatarID.ToOptionalUUID(),
			URL:     fmt.Sprintf(fileURLTemplate, user.AvatarID.ToOptionalUUID().String()+".jpg"),
		}
	}

	return query.UserData{
		UserID:   uuid.UUID(user.ID),
		Login:    user.Login,
		Password: user.Password,
		Avatar:   avatar,
	}, nil
}

func (s *userQueryService) GetByLogin(login string) (query.UserData, error) {
	const sqlQuery = `SELECT id, login, password
		              FROM user
		              WHERE login = ?
		              LIMIT 1`

	var user sqlxQueryUser
	err := s.client.GetContext(context.Background(), &user, sqlQuery, login)
	if err == sql.ErrNoRows {
		return query.UserData{}, errors.WithStack(domain.ErrUserNotFound)
	} else if err != nil {
		return query.UserData{}, err
	}

	var avatar *query.ImageData
	if user.AvatarID.ToOptionalUUID() != nil {
		avatar = &query.ImageData{
			ImageID: *user.AvatarID.ToOptionalUUID(),
			URL:     *user.AvatarPath,
		}
	}

	return query.UserData{
		UserID:   uuid.UUID(user.ID),
		Login:    user.Login,
		Password: user.Password,
		Avatar:   avatar,
	}, nil
}
