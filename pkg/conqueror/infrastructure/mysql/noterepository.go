package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
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
	if err != nil {
		return errors.WithStack(err)
	}

	err = repo.removeTags(note.ID())
	if err != nil {
		return err
	}

	return repo.storeTags(note)
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

func (repo *noteRepository) removeTags(noteID domain.NoteID) error {
	const sqlQuery = `DELETE FROM note_has_tag
		              WHERE note_id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(noteID))
	return errors.WithStack(err)

}

func (repo *noteRepository) storeTags(note *domain.Note) error {
	if len(note.Tags()) == 0 {
		return nil
	}

	const sqlQuery = `INSERT INTO note_has_tag (note_id, tag_id)
		              VALUES %s
		              ON DUPLICATE KEY UPDATE note_id=VALUES(note_id), tag_id=VALUES(tag_id)`

	queryPacks := make([]string, 0, len(note.Tags()))
	params := make([]interface{}, 0, len(note.Tags())*2)
	for _, tagID := range note.Tags() {
		queryPacks = append(queryPacks, "(?, ?)")
		params = append(params, binaryUUID(note.ID()), binaryUUID(tagID))
	}

	_, err := repo.client.ExecContext(repo.ctx, fmt.Sprintf(sqlQuery, strings.Join(queryPacks, ",")), params...)
	return errors.WithStack(err)
}
