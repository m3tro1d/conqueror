package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
	"github.com/pkg/errors"
)

func NewTaskTagRepository(ctx context.Context, client ClientContext) domain.TaskTagRepository {
	return &taskTagRepository{
		ctx:    ctx,
		client: client,
	}
}

type taskTagRepository struct {
	ctx    context.Context
	client ClientContext
}

func (repo *taskTagRepository) NextID() domain.TaskTagID {
	return domain.TaskTagID(uuid.Generate())
}

func (repo *taskTagRepository) Store(taskTag *domain.TaskTag) error {
	const sqlQuery = `INSERT INTO task_tag (id, name, user_id)
		              VALUES (?, ?, ?)
		              ON DUPLICATE KEY UPDATE user_id=VALUES(user_id)`

	args := []interface{}{
		binaryUUID(taskTag.ID()),
		taskTag.Name(),
		binaryUUID(taskTag.UserID()),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *taskTagRepository) GetByID(id domain.TaskTagID) (*domain.TaskTag, error) {
	const sqlQuery = `SELECT id, name, user_id
		              FROM task_tag
		              WHERE id = ?
		              LIMIT 1`

	var taskTag sqlxTaskTag
	err := repo.client.GetContext(repo.ctx, &taskTag, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrTaskTagNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewTaskTag(
		domain.TaskTagID(taskTag.ID),
		taskTag.Name,
		domain.UserID(taskTag.UserID),
	)
}

func (repo *taskTagRepository) RemoveByID(id domain.TaskTagID) error {
	const sqlQuery = `DELETE FROM task_tag
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}
