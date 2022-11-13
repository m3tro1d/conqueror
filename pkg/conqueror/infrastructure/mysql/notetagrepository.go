package mysql

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/jmoiron/sqlx"
)

func NewNoteTagRepository(ctx context.Context, client *sqlx.Conn) domain.NoteTagRepository {
	return &noteTagRepository{
		ctx:    ctx,
		client: client,
	}
}

type noteTagRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *noteTagRepository) NextID() domain.NoteTagID {
	return domain.NoteTagID(uuid.Generate())
}

func (repo *noteTagRepository) Store(noteTag *domain.NoteTag) error {
	const sqlQuery = `INSERT INTO note_tag (id, name, user_id)
		              VALUES (?, ?, ?)
		              ON DUPLICATE KEY UPDATE user_id=VALUES(user_id)`

	args := []interface{}{
		binaryUUID(noteTag.ID()),
		noteTag.Name(),
		binaryUUID(noteTag.UserID()),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *noteTagRepository) GetByID(id domain.NoteTagID) (*domain.NoteTag, error) {
	const sqlQuery = `SELECT id, name, user_id
		              FROM note_tag
		              WHERE id = ?
		              LIMIT 1`

	var noteTag sqlxNoteTag
	err := repo.client.GetContext(repo.ctx, &noteTag, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrNoteTagNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewNoteTag(
		domain.NoteTagID(noteTag.ID),
		noteTag.Name,
		domain.UserID(noteTag.UserID),
	)
}

func (repo *noteTagRepository) RemoveByID(id domain.NoteTagID) error {
	const sqlQuery = `DELETE FROM note_tag
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}
