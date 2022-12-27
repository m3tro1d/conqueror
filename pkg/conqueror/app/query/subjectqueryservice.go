package query

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
)

type SubjectQueryService interface {
	ListSubjects(ctx auth.UserContext) ([]SubjectData, error)
}

type SubjectData struct {
	ID    uuid.UUID
	Title string
}
