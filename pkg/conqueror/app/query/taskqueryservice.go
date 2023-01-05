package query

import (
	"time"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
)

type TaskQueryService interface {
	ListTasks(ctx auth.UserContext, spec ListTasksSpecification) ([]TaskData, error)
	GetTask(ctx auth.UserContext, taskID uuid.UUID) (TaskData, error)
}

type ListTasksSpecification struct {
	Query         string
	Sort          *TasksSortSettings
	ShowCompleted bool
	ForToday      bool
}

type TasksSortSettings struct {
	Field TasksSortField
	Order SortOrder
}

type TasksSortField int

const (
	TasksSortFieldStatus = TasksSortField(iota)
	TasksSortFieldTitle
)

type TaskData struct {
	ID           uuid.UUID
	DueDate      time.Time
	Title        string
	Description  string
	Status       TaskStatus
	SubjectID    *uuid.UUID
	SubjectTitle *string
}

type TaskStatus int

const (
	TaskStatusOpen = TaskStatus(iota)
	TaskStatusCompleted
)
