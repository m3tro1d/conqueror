package domain

import stderrors "errors"

var (
	ErrInvalidTimetableType = stderrors.New("invalid timetable type")
	ErrTimetableNotFound    = stderrors.New("timetable not found")
)

func NewTimetable(
	id TimetableID,
	userID UserID,
	timetableType TimetableType,
) (*Timetable, error) {
	return &Timetable{
		id:            id,
		userID:        userID,
		timetableType: timetableType,
	}, nil
}

type Timetable struct {
	id            TimetableID
	userID        UserID
	timetableType TimetableType
}

type TimetableType = int

const (
	TimetableTypeOneWeek = TimetableType(iota)
	TimetableTypeTwoWeeks
)

type TimetableRepository interface {
	NextID() TimetableID
	Store(timetable *Timetable) error
	GetByID(id TimetableID) (*Timetable, error)
	RemoveByID(id TimetableID) error
}

func (t *Timetable) ID() TimetableID {
	return t.id
}

func (t *Timetable) UserID() UserID {
	return t.userID
}

func (t *Timetable) TimetableType() TimetableType {
	return t.timetableType
}
