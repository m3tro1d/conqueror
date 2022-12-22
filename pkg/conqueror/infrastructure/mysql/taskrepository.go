package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/pkg/errors"
)

const (
	taskStatusOpen      = 0
	taskStatusCompleted = 1
)

func NewTaskRepository(ctx context.Context, client ClientContext) domain.TaskRepository {
	return &taskRepository{
		ctx:    ctx,
		client: client,
	}
}

type taskRepository struct {
	ctx    context.Context
	client ClientContext
}

func (repo *taskRepository) NextID() domain.TaskID {
	return domain.TaskID(uuid.Generate())
}

func (repo *taskRepository) Store(task *domain.Task) error {
	const sqlQuery = `INSERT INTO task (id, user_id, due_date, title, description, status, subject_id)
		              VALUES (?, ?, ?, ?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE user_id=VALUES(user_id), due_date=VALUES(due_date),
		                                      title=VALUES(title), description=VALUES(description),
		                                      status=VALUES(status), subject_id=VALUES(subject_id)`

	status, err := domainToDbTaskStatus(task.Status())
	if err != nil {
		return err
	}

	args := []interface{}{
		binaryUUID(task.ID()),
		binaryUUID(task.UserID()),
		task.DueDate(),
		task.Title(),
		task.Description(),
		status,
		makeNullBinaryUUID((*uuid.UUID)(task.SubjectID())),
	}

	_, err = repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	err = repo.removeTags(task.ID())
	if err != nil {
		return err
	}

	return repo.storeTags(task)
}

func (repo *taskRepository) GetByID(id domain.TaskID) (*domain.Task, error) {
	const sqlQuery = `SELECT id, user_id, due_date, title, description, status, subject_id
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

	status, err := dbToDomainTaskStatus(task.Status)
	if err != nil {
		return nil, err
	}

	return domain.NewTask(
		domain.TaskID(task.ID),
		domain.UserID(task.UserID),
		task.DueDate,
		task.Title,
		task.Description,
		status,
		(*domain.SubjectID)(task.SubjectID.ToOptionalUUID()),
	)
}

func (repo *taskRepository) RemoveByID(id domain.TaskID) error {
	const sqlQuery = `DELETE FROM task
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}

func (repo *taskRepository) removeTags(taskID domain.TaskID) error {
	const sqlQuery = `DELETE FROM task_has_tag
		              WHERE task_id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(taskID))
	return errors.WithStack(err)
}

func (repo *taskRepository) storeTags(task *domain.Task) error {
	if len(task.Tags()) == 0 {
		return nil
	}

	const sqlQuery = `INSERT INTO task_has_tag (task_id, tag_id)
		              VALUES %s
		              ON DUPLICATE KEY UPDATE task_id=VALUES(task_id), tag_id=VALUES(tag_id)`

	queryPacks := make([]string, 0, len(task.Tags()))
	params := make([]interface{}, 0, len(task.Tags())*2)
	for _, tagID := range task.Tags() {
		queryPacks = append(queryPacks, "(?, ?)")
		params = append(params, binaryUUID(task.ID()), binaryUUID(tagID))
	}

	_, err := repo.client.ExecContext(repo.ctx, fmt.Sprintf(sqlQuery, strings.Join(queryPacks, ",")), params...)
	return errors.WithStack(err)
}

func dbToDomainTaskStatus(status int) (domain.TaskStatus, error) {
	switch status {
	case taskStatusOpen:
		return domain.TaskStatusOpen, nil
	case taskStatusCompleted:
		return domain.TaskStatusCompleted, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidTaskStatus)
	}
}

func domainToDbTaskStatus(status domain.TaskStatus) (int, error) {
	switch status {
	case domain.TaskStatusOpen:
		return taskStatusOpen, nil
	case domain.TaskStatusCompleted:
		return taskStatusCompleted, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidTaskStatus)
	}
}
