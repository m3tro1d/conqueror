package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/app/query"
)

func NewTaskQueryService(client *sqlx.Conn) query.TaskQueryService {
	return &taskQueryService{
		client: client,
	}
}

type taskQueryService struct {
	client *sqlx.Conn
}

func (s *taskQueryService) ListTasks(ctx auth.UserContext) ([]query.TaskData, error) {
	const sqlQuery = `SELECT id, due_date, title, description, subject_id
		              FROM task
		              WHERE user_id = ?
		              ORDER BY due_date DESC`

	var tasks []sqlxQueryTask
	err := s.client.SelectContext(ctx, &tasks, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	taskIDToSqlxTagMap, err := s.getTasksTags(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]query.TaskData, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, query.TaskData{
			ID:          uuid.UUID(task.ID),
			DueDate:     task.DueDate,
			Title:       task.Title,
			Description: task.Description,
			Tags:        taskIDToSqlxTagMap[task.ID],
			SubjectID:   task.SubjectID.ToOptionalUUID(),
		})
	}

	return result, nil
}

func (s *taskQueryService) getTasksTags(ctx auth.UserContext) (map[binaryUUID][]query.TaskTagData, error) {
	const sqlQuery = `SELECT tag.id, task.id AS task_id, tag.name
				      FROM task_tag tag
				      INNER JOIN task ON task.id = tag.task_id
				      WHERE task.user_id = ?`

	var tags []sqlxQueryTaskTag
	err := s.client.SelectContext(ctx, &tags, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make(map[binaryUUID][]query.TaskTagData)
	for _, tag := range tags {
		result[tag.TaskID] = append(result[tag.TaskID], query.TaskTagData{
			ID:   uuid.UUID(tag.ID),
			Name: tag.Name,
		})
	}

	return result, nil
}
