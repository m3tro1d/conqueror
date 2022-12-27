package mysql

import (
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/app/query"

	"github.com/pkg/errors"
)

func NewSubjectQueryService(client ClientContext) query.SubjectQueryService {
	return &subjectQueryService{
		client: client,
	}
}

type subjectQueryService struct {
	client ClientContext
}

func (s *subjectQueryService) ListSubjects(ctx auth.UserContext) ([]query.SubjectData, error) {
	const sqlQuery = `SELECT id, title
		              FROM subject
		              WHERE user_id = ?
		              ORDER BY title`

	var subjects []sqlxQuerySubject
	err := s.client.SelectContext(ctx, &subjects, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]query.SubjectData, 0, len(subjects))
	for _, subject := range subjects {
		result = append(result, query.SubjectData{
			ID:    uuid.UUID(subject.ID),
			Title: subject.Title,
		})
	}

	return result, nil
}
