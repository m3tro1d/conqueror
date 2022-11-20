package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func NewScheduleRepository(ctx context.Context, client *sqlx.Conn) domain.ScheduleRepository {
	return &scheduleRepository{
		ctx:    ctx,
		client: client,
	}
}

type scheduleRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *scheduleRepository) NextID() domain.ScheduleID {
	return domain.ScheduleID(uuid.Generate())
}

func (repo *scheduleRepository) Store(schedule *domain.Schedule) error {
	const sqlQuery = `INSERT INTO schedule (id, timetable_id, is_even, title)
		              VALUES (?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE is_even=VALUES(is_even), title=VALUES(title)`

	args := []interface{}{
		binaryUUID(schedule.ID()),
		binaryUUID(schedule.TimetableID()),
		schedule.IsEven(),
		schedule.Title(),
	}

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *scheduleRepository) GetByID(id domain.ScheduleID) (*domain.Schedule, error) {
	const sqlQuery = `SELECT id, timetable_id, is_even, title
		              FROM schedule
		              WHERE id = ?
		              LIMIT 1`

	var schedule sqlxSchedule
	err := repo.client.GetContext(repo.ctx, &schedule, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrScheduleNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewSchedule(
		domain.ScheduleID(schedule.ID),
		domain.TimetableID(schedule.TimetableID),
		schedule.IsEven,
		schedule.Title,
	)
}

func (repo *scheduleRepository) RemoveByID(id domain.ScheduleID) error {
	const sqlQuery = `DELETE FROM schedule
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}
