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

func NewSchedule(
	id ScheduleID,
	timetableID TimetableID,
	isEven bool,
	title string,
) (*Schedule, error) {
	err := validateScheduleTitle(title)
	if err != nil {
		return nil, err
	}

	return &Schedule{
		id:          id,
		timetableID: timetableID,
		isEven:      isEven,
		title:       title,
	}, nil
}

type Schedule struct {
	id          ScheduleID
	timetableID TimetableID
	isEven      bool
	title       string
}

type ScheduleRepository interface {
	NextID() ScheduleID
	Store(schedule *Schedule) error
	GetByID(id ScheduleID) (*Schedule, error)
	RemoveByID(id ScheduleID) error
}

func (s *Schedule) ID() ScheduleID {
	return s.id
}

func (s *Schedule) TimetableID() TimetableID {
	return s.timetableID
}

func (s *Schedule) IsEven() bool {
	return s.isEven
}

func (s *Schedule) Title() string {
	return s.title
}

func validateScheduleTitle(title string) error {
	length := utf8.RuneCountInString(title)
	if length < minScheduleTitleLength && length > maxScheduleTitleLength {
		return errors.WithStack(ErrScheduleTitleLength)
	}
	return nil
}
