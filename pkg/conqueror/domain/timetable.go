package domain

import stderrors "errors"

var (
	ErrTimetableNotFound = stderrors.New("timetable not found")
)

func NewTimetable(
	id TimetableID,
	userID UserID,
	timetableType TimetableType,
	oddScheduleID ScheduleID,
	evenScheduleID *ScheduleID,
) (*Timetable, error) {
	return &Timetable{
		id:             id,
		userID:         userID,
		timetableType:  timetableType,
		oddScheduleID:  oddScheduleID,
		evenScheduleID: evenScheduleID,
	}, nil
}

type Timetable struct {
	id             TimetableID
	userID         UserID
	timetableType  TimetableType
	oddScheduleID  ScheduleID
	evenScheduleID *ScheduleID
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

func (t *Timetable) OddScheduleID() ScheduleID {
	return t.oddScheduleID
}

func (t *Timetable) EvenScheduleID() *ScheduleID {
	return t.evenScheduleID
}

func (t *Timetable) MakeOneWeek() error {
	t.timetableType = TimetableTypeOneWeek
	t.evenScheduleID = nil
	return nil
}

func (t *Timetable) MakeTwoWeeks(evenScheduleID ScheduleID) error {
	t.timetableType = TimetableTypeTwoWeeks
	t.evenScheduleID = &evenScheduleID
	return nil
}
