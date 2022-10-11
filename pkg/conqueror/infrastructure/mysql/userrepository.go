package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewUserRepository(ctx context.Context, client *sqlx.Conn) domain.UserRepository {
	return &userRepository{
		ctx:    ctx,
		client: client,
	}
}

type userRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *userRepository) NextID() domain.UserID {
	return domain.UserID(uuid.Generate())
}

func (repo *userRepository) Store(user *domain.User) error {
	const sqlQuery = `INSERT INTO user (id, login, password, nickname)
		              VALUES (?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE login=VALUES(login), password=VALUES(password),
		                                      nickname=VALUES(nickname)`

	args := []interface{}{
		user.ID(),
		user.Login(),
		user.Password(),
		user.Nickname(),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *userRepository) GetByID(id domain.UserID) (*domain.User, error) {
	const sqlQuery = `SELECT id, login, password, nickname
		              FROM user
		              WHERE id = ?`

	var user sqlxUser
	err := repo.client.SelectContext(repo.ctx, &user, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrUserNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewUser(
		domain.UserID(user.ID),
		user.Login,
		user.Password,
		user.Nickname,
	)
}

func (repo *userRepository) FindByLogin(login string) (*domain.User, error) {
	const sqlQuery = `SELECT id, login, password, nickname
		              FROM user
		              WHERE login = ?`

	var user sqlxUser
	err := repo.client.SelectContext(repo.ctx, &user, sqlQuery, login)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewUser(
		domain.UserID(user.ID),
		user.Login,
		user.Password,
		user.Nickname,
	)
}
