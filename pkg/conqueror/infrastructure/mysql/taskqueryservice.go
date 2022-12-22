package mysql

import (
	"database/sql"

	"conqueror/pkg/conqueror/domain"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/app/query"
)

func NewTaskQueryService(client ClientContext) query.TaskQueryService {
	return &taskQueryService{
		client: client,
	}
}

type taskQueryService struct {
	client ClientContext
}

func (s *taskQueryService) ListTasks(ctx auth.UserContext) ([]query.TaskData, error) {
	const sqlQuery = `SELECT id, due_date, title, description, status, subject_id
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
		status, err := dbToQueryTaskStatus(task.Status)
		if err != nil {
			return nil, err
		}

		result = append(result, query.TaskData{
			ID:          uuid.UUID(task.ID),
			DueDate:     task.DueDate,
			Title:       task.Title,
			Description: task.Description,
			Status:      status,
			Tags:        taskIDToSqlxTagMap[task.ID],
			SubjectID:   task.SubjectID.ToOptionalUUID(),
		})
	}

	return result, nil
}

func (s *taskQueryService) ListTaskTags(ctx auth.UserContext) ([]query.TaskTagData, error) {
	const sqlQuery = `SELECT id, name
		              FROM task_tag
		              WHERE user_id = ?
		              ORDER BY name`

	var tags []sqlxQueryTaskTag
	err := s.client.SelectContext(ctx, &tags, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]query.TaskTagData, 0, len(tags))
	for _, tag := range tags {
		result = append(result, query.TaskTagData{
			ID:   uuid.UUID(tag.ID),
			Name: tag.Name,
		})
	}

	return result, nil
}

func (s *taskQueryService) getTasksTags(ctx auth.UserContext) (map[binaryUUID][]query.TaskTagData, error) {
	const sqlQuery = `SELECT tag.id, t.id AS task_id, tag.name
				      FROM task_tag tag
				          INNER JOIN task_has_tag tht ON tag.id = tht.tag_id
				          INNER JOIN task t ON tht.task_id = t.id
				      WHERE t.user_id = ?`

	var tags []sqlxQueryTaskTagWithTask
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

func dbToQueryTaskStatus(status int) (query.TaskStatus, error) {
	switch status {
	case taskStatusOpen:
		return query.TaskStatusOpen, nil
	case taskStatusCompleted:
		return query.TaskStatusCompleted, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidTaskStatus)
	}
}