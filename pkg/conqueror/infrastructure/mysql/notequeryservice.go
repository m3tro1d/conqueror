package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"conqueror/pkg/conqueror/app/query"
)

func NewNoteQueryService(client *sqlx.Conn) query.NoteQueryService {
	return &noteQueryService{
		client: client,
	}
}

type noteQueryService struct {
	client *sqlx.Conn
}

func (s *noteQueryService) ListNotes(ctx auth.UserContext) ([]query.NoteData, error) {
	const sqlQuery = `SELECT id, due_date, title, description, subject_id
		              FROM note
		              WHERE user_id = ?
		              ORDER BY due_date DESC`

	var notes []sqlxQueryNote
	err := s.client.SelectContext(ctx, &notes, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	noteIDToSqlxTagMap, err := s.getNotesTags(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]query.NoteData, 0, len(notes))
	for _, note := range notes {
		result = append(result, query.NoteData{
			ID:        uuid.UUID(note.ID),
			Title:     note.Title,
			Content:   note.Content,
			Tags:      noteIDToSqlxTagMap[note.ID],
			SubjectID: note.SubjectID.ToOptionalUUID(),
		})
	}

	return result, nil
}

func (s *noteQueryService) getNotesTags(ctx auth.UserContext) (map[binaryUUID][]query.NoteTagData, error) {
	const sqlQuery = `SELECT tag.id, note.id AS note_id, tag.name
				      FROM note_tag tag
				      INNER JOIN note ON note.id = tag.note_id
				      WHERE note.user_id = ?`

	var tags []sqlxQueryNoteTag
	err := s.client.SelectContext(ctx, &tags, sqlQuery, binaryUUID(ctx.UserID()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make(map[binaryUUID][]query.NoteTagData)
	for _, tag := range tags {
		result[tag.NoteID] = append(result[tag.NoteID], query.NoteTagData{
			ID:   uuid.UUID(tag.ID),
			Name: tag.Name,
		})
	}

	return result, nil
}
