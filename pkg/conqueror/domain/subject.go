package domain

import (
	stderrors "errors"
	"fmt"
	"unicode/utf8"
)

const (
	minSubjectTitleLength = 1
	maxSubjectTitleLength = 100
)

var ErrSubjectTitleLength = fmt.Errorf("subject title must be greater or equal to %d and less or equal to %d", minSubjectTitleLength, maxSubjectTitleLength)
var ErrSubjectNotFound = stderrors.New("subject not found")

func NewSubject(id SubjectID, userID UserID, title string) (*Subject, error) {
	err := validateSubjectTitle(title)
	if err != nil {
		return nil, err
	}

	return &Subject{
		id:     id,
		userID: userID,
		title:  title,
	}, nil
}

type Subject struct {
	id     SubjectID
	userID UserID
	title  string
}

type SubjectRepository interface {
	NextID() SubjectID
	Store(subject *Subject) error
	GetByID(id SubjectID) (*Subject, error)
	RemoveByID(id SubjectID) error
}

func (s *Subject) ID() SubjectID {
	return s.id
}

func (s *Subject) UserID() UserID {
	return s.userID
}

func (s *Subject) Title() string {
	return s.title
}

func (s *Subject) ChangeTitle(newTitle string) error {
	err := validateSubjectTitle(newTitle)
	if err != nil {
		return err
	}

	s.title = newTitle
	return nil
}

func validateSubjectTitle(title string) error {
	length := utf8.RuneCountInString(title)
	if length < minSubjectTitleLength || length > maxSubjectTitleLength {
		return ErrSubjectTitleLength
	}
	return nil
}
