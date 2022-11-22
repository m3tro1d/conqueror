package mysql

import (
	"context"
	"database/sql"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	weekdayMonday    = 0
	weekdayTuesday   = 1
	weekdayWednesday = 2
	weekdayThursday  = 3
	weekdayFriday    = 4
	weekdaySaturday  = 5
	weekdaySunday    = 6
)

func NewLessonIntervalRepository(ctx context.Context, client *sqlx.Conn) domain.LessonIntervalRepository {
	return &lessonIntervalRepository{
		ctx:    ctx,
		client: client,
	}
}

type lessonIntervalRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *lessonIntervalRepository) NextID() domain.LessonIntervalID {
	return domain.LessonIntervalID(uuid.Generate())
}

func (repo *lessonIntervalRepository) Store(lessonInterval *domain.LessonInterval) error {
	const sqlQuery = `INSERT INTO lesson_interval (id, schedule_id, weekday, start_time, end_time)
		              VALUES (?, ?, ?, ?, ?)
		              ON DUPLICATE KEY UPDATE weekday=VALUES(weekday), start_time=VALUES(start_time),
		                                      end_time=VALUES(end_time)`

	weekday, err := domainToDbWeekday(lessonInterval.Weekday())
	if err != nil {
		return errors.WithStack(err)
	}

	args := []interface{}{
		binaryUUID(lessonInterval.ID()),
		binaryUUID(lessonInterval.ScheduleID()),
		weekday,
		lessonInterval.StartTime(),
		lessonInterval.EndTime(),
	}

	_, err = repo.client.ExecContext(repo.ctx, sqlQuery, args...)
	return errors.WithStack(err)
}

func (repo *lessonIntervalRepository) GetByID(id domain.LessonIntervalID) (*domain.LessonInterval, error) {
	const sqlQuery = `SELECT id, schedule_id, weekday, start_time, end_time
		              FROM lesson_interval
		              WHERE id = ?
		              LIMIT 1`

	var lessonInterval sqlxLessonInterval
	err := repo.client.GetContext(repo.ctx, &lessonInterval, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrLessonIntervalNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	weekday, err := dbToDomainWeekday(lessonInterval.Weekday)
	if err != nil {
		return nil, err
	}

	return domain.NewLessonInterval(
		domain.LessonIntervalID(lessonInterval.ID),
		domain.ScheduleID(lessonInterval.ScheduleID),
		weekday,
		lessonInterval.StartTime,
		lessonInterval.EndTime,
	)
}

func (repo *lessonIntervalRepository) RemoveByID(id domain.LessonIntervalID) error {
	const sqlQuery = `DELETE FROM lesson_interval
		              WHERE id = ?`

	_, err := repo.client.ExecContext(repo.ctx, sqlQuery, binaryUUID(id))
	return errors.WithStack(err)
}

func domainToDbWeekday(weekday domain.Weekday) (int, error) {
	switch weekday {
	case domain.WeekdayMonday:
		return weekdayMonday, nil
	case domain.WeekdayTuesday:
		return weekdayTuesday, nil
	case domain.WeekdayWednesday:
		return weekdayWednesday, nil
	case domain.WeekdayThursday:
		return weekdayThursday, nil
	case domain.WeekdayFriday:
		return weekdayFriday, nil
	case domain.WeekdaySaturday:
		return weekdaySaturday, nil
	case domain.WeekdaySunday:
		return weekdaySunday, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidWeekday)
	}
}

func dbToDomainWeekday(weekday int) (domain.Weekday, error) {
	switch weekday {
	case weekdayMonday:
		return domain.WeekdayMonday, nil
	case weekdayTuesday:
		return domain.WeekdayTuesday, nil
	case weekdayWednesday:
		return domain.WeekdayWednesday, nil
	case weekdayThursday:
		return domain.WeekdayThursday, nil
	case weekdayFriday:
		return domain.WeekdayFriday, nil
	case weekdaySaturday:
		return domain.WeekdaySaturday, nil
	case weekdaySunday:
		return domain.WeekdaySunday, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidWeekday)
	}
}
