package domain

import (
	stderrors "errors"
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	minScheduleTitleLength = 1
	maxScheduleTitleLength = 100
)

var (
	ErrScheduleTitleLength = fmt.Errorf("lesson auditorium must be greater than or equal to %d and less or equal to %d", minScheduleTitleLength, maxScheduleTitleLength)
	ErrScheduleNotFound    = stderrors.New("schedule not found")
)

func NewSchedule(title string, lessonIntervals map[Weekday]LessonIntervalID) (*Schedule, error) {
	err := validateScheduleTitle(title)
	if err != nil {
		return nil, err
	}

	return &Schedule{
		title:           title,
		lessonIntervals: lessonIntervals,
	}, nil
}

type Schedule struct {
	title           string
	lessonIntervals map[Weekday]LessonIntervalID
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

type ScheduleInterface interface {
	NextID() ScheduleID
	Store(schedule *Schedule) error
	GetByID(id ScheduleID) (*Schedule, error)
	RemoveByID(id ScheduleID) error
}

func validateScheduleTitle(title string) error {
	length := utf8.RuneCountInString(title)
	if length < minScheduleTitleLength && length > maxScheduleTitleLength {
		return errors.WithStack(ErrScheduleTitleLength)
	}
	return nil
}
