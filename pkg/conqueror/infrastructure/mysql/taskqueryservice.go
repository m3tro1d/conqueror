package mysql

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/domain"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func NewTaskQueryService(client ClientContext) query.TaskQueryService {
	return &taskQueryService{
		client: client,
	}
}

type taskQueryService struct {
	client ClientContext
}

func (s *taskQueryService) ListTasks(ctx auth.UserContext, spec query.ListTasksSpecification) ([]query.TaskData, error) {
	const sqlQuery = `SELECT t.id, t.due_date, t.title, t.description, t.status, t.subject_id, s.title AS subject_title
		              FROM task t
		              	LEFT JOIN subject s on t.subject_id = s.id
		              WHERE %s
		              ORDER BY %s`

	var whereClauses []string
	var args []interface{}

	whereClauses = append(whereClauses, "t.user_id = ?")
	args = append(args, binaryUUID(ctx.UserID()))
	if !spec.ShowCompleted {
		whereClauses = append(whereClauses, "t.status <> ?")
		args = append(args, taskStatusCompleted)
	}
	if spec.Query != "" {
		whereClauses = append(whereClauses, "t.title LIKE ?")
		args = append(args, "%"+spec.Query+"%")
	}
	if spec.ForToday {
		whereClauses = append(whereClauses, "t.due_date = CURDATE()")
	}

	var orders []string
	if spec.Sort != nil {
		sort := ""
		switch spec.Sort.Field {
		case query.TasksSortFieldStatus:
			sort += "t.status"
		case query.TasksSortFieldTitle:
			sort += "t.title"
		}

		switch spec.Sort.Order {
		case query.SortOrderAsc:
			sort += " ASC"
		case query.SortOrderDesc:
			sort += " DESC"
		}

		orders = append(orders, sort)
	}
	orders = append(orders, "t.due_date ASC")

	var tasks []sqlxQueryTask
	err := s.client.SelectContext(
		ctx,
		&tasks,
		fmt.Sprintf(
			sqlQuery,
			strings.Join(whereClauses, " AND "),
			strings.Join(orders, ", "),
		),
		args...,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]query.TaskData, 0, len(tasks))
	for _, task := range tasks {
		status, err := dbToQueryTaskStatus(task.Status)
		if err != nil {
			return nil, err
		}

		result = append(result, query.TaskData{
			ID:           uuid.UUID(task.ID),
			DueDate:      task.DueDate,
			Title:        task.Title,
			Description:  task.Description,
			Status:       status,
			SubjectID:    task.SubjectID.ToOptionalUUID(),
			SubjectTitle: task.SubjectTitle,
		})
	}

	return result, nil
}

func (s *taskQueryService) GetTask(ctx auth.UserContext, taskID uuid.UUID) (query.TaskData, error) {
	const sqlQuery = `SELECT t.id, t.due_date, t.title, t.description, t.status, t.subject_id, s.title AS subject_title
		              FROM task t
		              	LEFT JOIN subject s on t.subject_id = s.id
		              WHERE t.user_id = ? AND t.id = ?`

	var task sqlxTask
	err := s.client.GetContext(ctx, &task, sqlQuery, binaryUUID(ctx.UserID()), binaryUUID(taskID))
	if err == sql.ErrNoRows {
		return query.TaskData{}, errors.WithStack(domain.ErrTaskNotFound)
	} else if err != nil {
		return query.TaskData{}, errors.WithStack(err)
	}

	status, err := dbToQueryTaskStatus(task.Status)
	if err != nil {
		return query.TaskData{}, err
	}

	return query.TaskData{
		ID:           uuid.UUID(task.ID),
		DueDate:      task.DueDate,
		Title:        task.Title,
		Description:  task.Description,
		Status:       status,
		SubjectID:    task.SubjectID.ToOptionalUUID(),
		SubjectTitle: task.SubjectTitle,
	}, nil
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
