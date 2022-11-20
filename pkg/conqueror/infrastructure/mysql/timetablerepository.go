package mysql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

const (
	timetableTypeOneWeek  = 0
	timetableTypeTwoWeeks = 1
)

func NewTimetableRepository(ctx context.Context, client *sqlx.Conn) domain.TimetableRepository {
	return &timetableRepository{
		ctx:    ctx,
		client: client,
	}
}

type timetableRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *timetableRepository) NextID() domain.TimetableID {
	return domain.TimetableID(uuid.Generate())
}

func (repo *timetableRepository) Store(timetable *domain.Timetable) error {
	const sqlQuery = `INSERT INTO timetable (id, user_id, type)
		              VALUES (?, ?, ?)
		              ON DUPLICATE KEY UPDATE type=VALUES(type)`

	timetableType, err := domainToDbTimetableType(timetable.TimetableType())
	if err != nil {
		return err
	}

	args := []interface{}{
		binaryUUID(timetable.ID()),
		binaryUUID(timetable.UserID()),
		timetableType,
	}

	_, err = repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *timetableRepository) GetByID(id domain.TimetableID) (*domain.Timetable, error) {
	const sqlQuery = `SELECT id, user_id, type
		              FROM timetable
		              WHERE id = ?
		              LIMIT 1`

	var timetable sqlxTimetable
	err := repo.client.GetContext(repo.ctx, &timetable, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrTimetableNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	timetableType, err := dbToDomainTimetableType(timetable.TimetableType)
	if err != nil {
		return nil, err
	}

	return domain.NewTimetable(
		domain.TimetableID(timetable.ID),
		domain.UserID(timetable.UserID),
		timetableType,
	)
}

func (repo *timetableRepository) RemoveByID(id domain.TimetableID) error {
	const sqlQuery = `DELETE FROM timetable
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}

func domainToDbTimetableType(timetableType domain.TimetableType) (int, error) {
	switch timetableType {
	case domain.TimetableTypeOneWeek:
		return timetableTypeOneWeek, nil
	case domain.TimetableTypeTwoWeeks:
		return timetableTypeTwoWeeks, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidTimetableType)
	}
}

func dbToDomainTimetableType(timetableType int) (domain.TimetableType, error) {
	switch timetableType {
	case timetableTypeOneWeek:
		return domain.TimetableTypeOneWeek, nil
	case timetableTypeTwoWeeks:
		return domain.TimetableTypeTwoWeeks, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidTimetableType)
	}
}
