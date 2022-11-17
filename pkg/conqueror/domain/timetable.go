package domain

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
	oddSchedule   Schedule
	evenSchedule  *Schedule
}

type TimetableType = int

const (
	TimetableTypeOneWeek = TimetableType(iota)
	TimetableTypeTwoWeeks
)

type Schedule struct {
	title           string
	lessonIntervals map[Weekday]LessonInterval
}

type LessonInterval struct {
	startTime string
	endTime   string
	lesson    Lesson
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
