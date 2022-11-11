package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewNoteRepository(ctx context.Context, client *sqlx.Conn) domain.NoteRepository {
	return &noteRepository{
		ctx:    ctx,
		client: client,
	}
}

type noteRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *noteRepository) NextID() domain.NoteID {
	return domain.NoteID(uuid.Generate())
}

func (repo *noteRepository) Store(note *domain.Note) error {
	const sqlQuery = `INSERT INTO note (id, user_id, title, content, subject_id)
		              VALUES (?, ?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE user_id=VALUES(user_id), title=VALUES(title),
		                                      content=VALUES(content), subject_id=VALUES(subject_id)`

	args := []interface{}{
		binaryUUID(note.ID()),
		binaryUUID(note.UserID()),
		note.Title(),
		note.Content(),
		makeNullBinaryUUID((*uuid.UUID)(note.SubjectID())),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *noteRepository) GetByID(id domain.NoteID) (*domain.Note, error) {
	const sqlQuery = `SELECT id, user_id, title, content, updated_at, subject_id
		              FROM note
		              WHERE id = ?
		              LIMIT 1`

	var note sqlxNote
	err := repo.client.GetContext(repo.ctx, &note, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrNoteNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewNote(
		domain.NoteID(note.ID),
		domain.UserID(note.UserID),
		note.Title,
		note.Content,
		(*domain.SubjectID)(note.SubjectID.ToOptionalUUID()),
	)
}

func (repo *noteRepository) RemoveByID(id domain.NoteID) error {
	const sqlQuery = `DELETE FROM note
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}
