package query

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"time"
)

type NoteQueryService interface {
	ListNotes(ctx auth.UserContext) ([]NoteData, error)
	ListNoteTags(ctx auth.UserContext) ([]NoteTagData, error)
}

type NoteData struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Tags      []NoteTagData
	UpdatedAt time.Time
	SubjectID *uuid.UUID
}

type NoteTagData struct {
	ID   uuid.UUID
	Name string
}
