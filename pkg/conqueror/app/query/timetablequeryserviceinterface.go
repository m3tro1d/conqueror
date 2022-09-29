package query

import "time"

type TimetableQueryServiceInterface interface {
	GetTimetableForToday() (DayTimetable, error)
	GetTimetableBetweenDays(startDate, endDate time.Time) ([]DayTimetable, error)
}

type DayTimetable struct {
	Lessons []Lesson
}

type Lesson struct {
	Interval   TimeInterval
	Title      string
	Auditorium string
}

type TimeInterval struct {
	StartTime string
	EndTime   string
}
