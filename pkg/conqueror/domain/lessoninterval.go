package domain

import stderrors "errors"

var (
	ErrInvalidWeekday         = stderrors.New("invalid weekday")
	ErrLessonIntervalNotFound = stderrors.New("lesson interval not found")
)

func NewLessonInterval(
	id LessonIntervalID,
	scheduleID ScheduleID,
	weekday Weekday,
	startTime, endTime string,
) (*LessonInterval, error) {
	return &LessonInterval{
		id:         id,
		scheduleID: scheduleID,
		weekday:    weekday,
		startTime:  startTime,
		endTime:    endTime,
	}, nil
}

type LessonInterval struct {
	id         LessonIntervalID
	scheduleID ScheduleID
	weekday    Weekday
	startTime  string
	endTime    string
}

type Weekday int

const (
	WeekdayMonday = Weekday(iota)
	WeekdayTuesday
	WeekdayWednesday
	WeekdayThursday
	WeekdayFriday
	WeekdaySaturday
	WeekdaySunday
)

type LessonIntervalRepository interface {
	NextID() LessonIntervalID
	Store(schedule *LessonInterval) error
	GetByID(id LessonIntervalID) (*LessonInterval, error)
	RemoveByID(id LessonIntervalID) error
}

func (i *LessonInterval) ID() LessonIntervalID {
	return i.id
}

func (i *LessonInterval) ScheduleID() ScheduleID {
	return i.scheduleID
}

func (i *LessonInterval) Weekday() Weekday {
	return i.weekday
}

func (i *LessonInterval) StartTime() string {
	return i.startTime
}

func (i *LessonInterval) EndTime() string {
	return i.endTime
}
