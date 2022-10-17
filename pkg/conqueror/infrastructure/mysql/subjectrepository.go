package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewSubjectRepository(ctx context.Context, client *sqlx.Conn) domain.SubjectRepository {
	return &subjectRepository{
		ctx:    ctx,
		client: client,
	}
}

type subjectRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *subjectRepository) NextID() domain.SubjectID {
	return domain.SubjectID(uuid.Generate())
}

func (repo *subjectRepository) Store(subject *domain.Subject) error {
	const sqlQuery = `INSERT INTO subject (id, user_id, title)
		              VALUES (?, ?, ?)
		              ON DUPLICATE KEY UPDATE user_id=VALUES(user_id), title=VALUES(title)`

	args := []interface{}{
		binaryUUID(subject.ID()),
		binaryUUID(subject.UserID()),
		subject.Title(),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *subjectRepository) GetByID(id domain.SubjectID) (*domain.Subject, error) {
	const sqlQuery = `SELECT id, user_id, title
		              FROM subject
		              WHERE id = ?
		              LIMIT 1`

	var subject sqlxSubject
	err := repo.client.GetContext(repo.ctx, &subject, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrSubjectNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewSubject(
		domain.SubjectID(subject.ID),
		domain.UserID(subject.UserID),
		subject.Title,
	)
}

func (repo *subjectRepository) RemoveByID(id domain.SubjectID) error {
	const sqlQuery = `DELETE FROM subject
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}
