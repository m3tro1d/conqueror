package query

import (
	"time"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
)

type TaskQueryService interface {
	ListTasks(ctx auth.UserContext) ([]TaskData, error)
	ListTaskTags(ctx auth.UserContext) ([]TaskTagData, error)
}

type TaskData struct {
	ID          uuid.UUID
	DueDate     time.Time
	Title       string
	Description string
	Tags        []TaskTagData
	SubjectID   *uuid.UUID
}

type TaskTagData struct {
	ID   uuid.UUID
	Name string
}
