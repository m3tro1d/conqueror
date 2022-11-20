package domain

func NewLessonInterval(startTime, endTime string, scheduleID ScheduleID) *LessonInterval {
	return &LessonInterval{
		startTime:  startTime,
		endTime:    endTime,
		scheduleID: scheduleID,
	}
}

type LessonInterval struct {
	startTime  string
	endTime    string
	scheduleID ScheduleID
}
