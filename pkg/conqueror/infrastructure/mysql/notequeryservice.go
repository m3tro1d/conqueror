package mysql

import (
	"conqueror/pkg/conqueror/domain"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/app/query"
)

func NewNoteQueryService(client ClientContext) query.NoteQueryService {
	return &noteQueryService{
		client: client,
	}
}

type noteQueryService struct {
	client ClientContext
}

func (s *noteQueryService) ListNotes(ctx auth.UserContext, spec query.ListNotesSpecification) ([]query.NoteData, error) {
	const sqlQuery = `SELECT id, title, content, updated_at, subject_id
		              FROM note
		              WHERE %s
		              ORDER BY updated_at DESC`

	var whereClauses []string
	var args []interface{}

	whereClauses = append(whereClauses, "user_id = ?")
	args = append(args, binaryUUID(ctx.UserID()))
	if spec.Query != "" {
		whereClauses = append(whereClauses, "title LIKE ?")
		args = append(args, "%"+spec.Query+"%")
	}

	var notes []sqlxQueryNote
	err := s.client.SelectContext(
		ctx,
		&notes,
		fmt.Sprintf(sqlQuery, strings.Join(whereClauses, " AND ")),
		args...,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]query.NoteData, 0, len(notes))
	for _, note := range notes {
		result = append(result, query.NoteData{
			ID:        uuid.UUID(note.ID),
			Title:     note.Title,
			Content:   note.Content,
			UpdatedAt: note.UpdatedAt,
			SubjectID: note.SubjectID.ToOptionalUUID(),
		})
	}

	return result, nil
}

func (s *noteQueryService) GetNote(ctx auth.UserContext, noteID uuid.UUID) (query.NoteData, error) {
	const sqlQuery = `SELECT n.id, n.title, n.content, n.updated_at, subject_id, s.title AS subject_title
		              FROM note n
		              	LEFT JOIN subject s on n.subject_id = s.id
		              WHERE n.user_id = ? AND n.id = ?`

	var note sqlxNote
	err := s.client.GetContext(ctx, &note, sqlQuery, binaryUUID(ctx.UserID()), binaryUUID(noteID))
	if err == sql.ErrNoRows {
		return query.NoteData{}, errors.WithStack(domain.ErrNoteNotFound)
	} else if err != nil {
		return query.NoteData{}, errors.WithStack(err)
	}

	return query.NoteData{
		ID:           uuid.UUID(note.ID),
		Title:        note.Title,
		Content:      note.Content,
		UpdatedAt:    note.UpdatedAt,
		SubjectID:    note.SubjectID.ToOptionalUUID(),
		SubjectTitle: note.SubjectTitle,
	}, nil
}
