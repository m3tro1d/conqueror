package domain

import (
	stderrors "errors"
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	maxLessonAuditoriumLength = 50
)

var (
	ErrLessonAuditoriumLength = fmt.Errorf("lesson auditorium must be less or equal to %d", maxLessonAuditoriumLength)
	ErrLessonNotFound         = stderrors.New("lesson not found")
)

func NewLesson(id LessonID, lessonType LessonType, auditorium string) (*Lesson, error) {
	err := validateLessonAuditorium(auditorium)
	if err != nil {
		return nil, err
	}

	return &Lesson{
		id:         id,
		lessonType: lessonType,
		auditorium: auditorium,
	}, nil
}

type Lesson struct {
	id         LessonID
	lessonType LessonType
	auditorium string
}

type LessonType int

const (
	LessonTypeLecture = LessonType(iota)
	LessonTypePractice
	LessonTypeLab
)

type LessonRepository interface {
	NextID() LessonID
	Store(lesson *Lesson) error
	GetByID(id LessonID) (*Lesson, error)
	RemoveByID(id LessonID) error
}

func (l *Lesson) ID() LessonID {
	return l.id
}

func (l *Lesson) LessonType() LessonType {
	return l.lessonType
}

func (l *Lesson) Auditorium() string {
	return l.auditorium
}

func (l *Lesson) ChangeLessonType(newLessonType LessonType) error {
	l.lessonType = newLessonType
	return nil
}

func (l *Lesson) ChangeAuditorium(newAuditorium string) error {
	err := validateLessonAuditorium(newAuditorium)
	if err != nil {
		return err
	}

	l.auditorium = newAuditorium
	return nil
}

func validateLessonAuditorium(auditorium string) error {
	length := utf8.RuneCountInString(auditorium)
	if length > maxLessonAuditoriumLength {
		return errors.WithStack(ErrLessonAuditoriumLength)
	}
	return nil
}
