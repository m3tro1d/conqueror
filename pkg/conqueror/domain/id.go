package domain

import "conqueror/pkg/common/uuid"

type (
	UserID    uuid.UUID
	SubjectID uuid.UUID

	TimetableID      uuid.UUID
	ScheduleID       uuid.UUID
	LessonIntervalID uuid.UUID
	LessonID         uuid.UUID

	TaskID    uuid.UUID
	TaskTagID uuid.UUID

	NoteID    uuid.UUID
	NoteTagID uuid.UUID
)
