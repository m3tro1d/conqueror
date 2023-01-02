package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/pkg/errors"
)

func NewUserRepository(ctx context.Context, client ClientContext) domain.UserRepository {
	return &userRepository{
		ctx:    ctx,
		client: client,
	}
}

type userRepository struct {
	ctx    context.Context
	client ClientContext
}

func (repo *userRepository) NextID() domain.UserID {
	return domain.UserID(uuid.Generate())
}

func (repo *userRepository) Store(user *domain.User) error {
	const sqlQuery = `INSERT INTO user (id, login, password, avatar_id)
		              VALUES (?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE login=VALUES(login), password=VALUES(password),
		                                      avatar_id=VALUES(avatar_id)`

	args := []interface{}{
		binaryUUID(user.ID()),
		user.Login(),
		user.Password(),
		makeNullBinaryUUID((*uuid.UUID)(user.AvatarID())),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *userRepository) GetByID(id domain.UserID) (*domain.User, error) {
	const sqlQuery = `SELECT id, login, password, avatar_id
		              FROM user
		              WHERE id = ?
		              LIMIT 1`

	var user sqlxUser
	err := repo.client.GetContext(repo.ctx, &user, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrUserNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewUser(
		domain.UserID(user.ID),
		user.Login,
		user.Password,
		(*domain.ImageID)(user.AvatarID.ToOptionalUUID()),
	)
}

func (repo *userRepository) FindByLogin(login string) (*domain.User, error) {
	const sqlQuery = `SELECT id, login, password, avatar_id
		              FROM user
		              WHERE login = ?
		              LIMIT 1`

	var user sqlxUser
	err := repo.client.GetContext(repo.ctx, &user, sqlQuery, login)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewUser(
		domain.UserID(user.ID),
		user.Login,
		user.Password,
		(*domain.ImageID)(user.AvatarID.ToOptionalUUID()),
	)
}
