package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

func NewTaskRepository(ctx context.Context, client *sqlx.Conn) domain.TaskRepository {
	return &taskRepository{
		ctx:    ctx,
		client: client,
	}
}

type taskRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *taskRepository) NextID() domain.TaskID {
	return domain.TaskID(uuid.Generate())
}

func (repo *taskRepository) Store(task *domain.Task) error {
	const sqlQuery = `INSERT INTO task (id, user_id, due_date, title, description)
		              VALUES (?, ?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE user_id=VALUES(user_id), due_date=VALUES(due_date),
		                                      title=VALUES(title), description=VALUES(description)`

	args := []interface{}{
		binaryUUID(task.ID()),
		binaryUUID(task.UserID()),
		task.DueDate(),
		task.Title(),
		task.Description(),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *taskRepository) GetByID(id domain.TaskID) (*domain.Task, error) {
	const sqlQuery = `SELECT id, user_id, due_date, title, description
		              FROM task
		              WHERE id = ?
		              LIMIT 1`

	var task sqlxTask
	err := repo.client.GetContext(repo.ctx, &task, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrTaskNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewTask(
		domain.TaskID(task.ID),
		domain.UserID(task.UserID),
		task.DueDate,
		task.Title,
		task.Description,
	)
}

func (repo *taskRepository) RemoveByID(id domain.TaskID) error {
	const sqlQuery = `DELETE FROM task
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}
