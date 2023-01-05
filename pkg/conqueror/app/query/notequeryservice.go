package query

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/auth"
	"time"
)

type NoteQueryService interface {
	ListNotes(ctx auth.UserContext, spec ListNotesSpecification) ([]NoteData, error)
	GetNote(ctx auth.UserContext, noteID uuid.UUID) (NoteData, error)
}

type ListNotesSpecification struct {
	Query string
}

type NoteData struct {
	ID           uuid.UUID
	Title        string
	Content      string
	UpdatedAt    time.Time
	SubjectID    *uuid.UUID
	SubjectTitle *string
}
