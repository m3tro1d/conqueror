package domain

import (
	"fmt"
	"unicode/utf8"
)

const (
	minSubjectTitleLength = 1
	maxSubjectTitleLength = 100
)

var ErrSubjectTitleLength = fmt.Errorf("subject title must be more or equal to %d and less or equal to %d", minSubjectTitleLength, maxSubjectTitleLength)

func NewSubject(id SubjectID, userID UserID, title string) *Subject {
	return &Subject{
		id:     id,
		userID: userID,
		title:  title,
	}
}

type Subject struct {
	id     SubjectID
	userID UserID
	title  string
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
