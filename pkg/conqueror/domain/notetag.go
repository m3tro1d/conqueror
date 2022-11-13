package domain

import (
	stderrors "errors"
	"fmt"
	"unicode/utf8"
)

const (
	minNoteTagNameLength = 1
	maxNoteTagNameLength = 200
)

var (
	ErrNoteTagNameLength = fmt.Errorf("note tag name must be greater or equal to %d and less or equal to %d", minNoteTagNameLength, maxNoteTagNameLength)
	ErrNoteTagNotFound   = stderrors.New("note tag not found")
)

func NewNoteTag(id NoteTagID, name string, userID UserID) (*NoteTag, error) {
	err := validateNoteTagName(name)
	if err != nil {
		return nil, err
	}

	return &NoteTag{
		id:     id,
		name:   name,
		userID: userID,
	}, nil
}

type NoteTag struct {
	id     NoteTagID
	name   string
	userID UserID
}

type NoteTagRepository interface {
	NextID() NoteTagID
	Store(noteTag *NoteTag) error
	GetByID(id NoteTagID) (*NoteTag, error)
	RemoveByID(id NoteTagID) error
}

func (t *NoteTag) ID() NoteTagID {
	return t.id
}

func (t *NoteTag) Name() string {
	return t.name
}

func (t *NoteTag) UserID() UserID {
	return t.userID
}

func (t *NoteTag) ChangeName(newName string) error {
	err := validateNoteTagName(newName)
	if err != nil {
		return err
	}

	t.name = newName
	return nil
}

func validateNoteTagName(name string) error {
	length := utf8.RuneCountInString(name)
	if length < minNoteTagNameLength || length > maxNoteTagNameLength {
		return ErrNoteTagNameLength
	}
	return nil
}
