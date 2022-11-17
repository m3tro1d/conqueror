package domain

import (
	stderrors "errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	minNoteTitleLength   = 1
	maxNoteTitleLength   = 200
	maxNoteContentLength = 2000
)

var (
	ErrNoteTitleLength   = fmt.Errorf("note title must be greater or equal to %d and less or equal to %d", minNoteTitleLength, maxNoteTitleLength)
	ErrNoteContentLength = fmt.Errorf("note content must be less or equal to %d", maxNoteContentLength)
	ErrDuplicateNoteTags = stderrors.New("duplicate note tags")
	ErrNoteNotFound      = stderrors.New("note not found")
)

func NewNote(id NoteID, userID UserID, title, content string, subjectID *SubjectID) (*Note, error) {
	err := validateNoteTitle(title)
	if err != nil {
		return nil, err
	}

	err = validateNoteContent(content)
	if err != nil {
		return nil, err
	}

	return &Note{
		id:        id,
		userID:    userID,
		title:     title,
		content:   content,
		updatedAt: time.Now(),
		subjectID: subjectID,
	}, nil
}

type Note struct {
	id        NoteID
	userID    UserID
	title     string
	content   string
	tags      []NoteTagID
	updatedAt time.Time
	subjectID *SubjectID
}

type NoteRepository interface {
	NextID() NoteID
	Store(note *Note) error
	GetByID(id NoteID) (*Note, error)
	RemoveByID(id NoteID) error
}

func (n *Note) ID() NoteID {
	return n.id
}

func (n *Note) UserID() UserID {
	return n.userID
}

func (n *Note) Title() string {
	return n.title
}

func (n *Note) Content() string {
	return n.content
}

func (n *Note) Tags() []NoteTagID {
	return n.tags
}

func (n *Note) UpdatedAt() time.Time {
	return n.updatedAt
}

func (n *Note) SubjectID() *SubjectID {
	return n.subjectID
}

func (n *Note) ChangeTitle(newTitle string) error {
	err := validateNoteTitle(newTitle)
	if err != nil {
		return err
	}

	n.title = newTitle
	return nil
}

func (n *Note) ChangeContent(newContent string) error {
	err := validateNoteContent(newContent)
	if err != nil {
		return err
	}

	n.content = newContent
	return nil
}

func (n *Note) ChangeTags(tags []NoteTagID) error {
	err := validateNoteTags(tags)
	if err != nil {
		return err
	}

	n.tags = tags
	return nil
}

func validateNoteTitle(title string) error {
	length := utf8.RuneCountInString(title)
	if length < minNoteTitleLength || length > maxNoteTitleLength {
		return ErrNoteTitleLength
	}
	return nil
}

func validateNoteContent(description string) error {
	length := utf8.RuneCountInString(description)
	if length > maxNoteContentLength {
		return ErrNoteContentLength
	}
	return nil
}

func validateNoteTags(tags []NoteTagID) error {
	tagsMap := make(map[NoteTagID]bool)
	for _, tagID := range tags {
		if tagsMap[tagID] {
			return errors.WithStack(ErrDuplicateNoteTags)
		}
		tagsMap[tagID] = true
	}
	return nil
}
