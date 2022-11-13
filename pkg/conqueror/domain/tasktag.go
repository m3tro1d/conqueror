package domain

import (
	stderrors "errors"
	"fmt"
	"unicode/utf8"
)

const (
	minTaskTagNameLength = 1
	maxTaskTagNameLength = 200
)

var (
	ErrTaskTagNameLength = fmt.Errorf("task tag name must be greater or equal to %d and less or equal to %d", minTaskTagNameLength, maxTaskTagNameLength)
	ErrTaskTagNotFound   = stderrors.New("task tag not found")
)

func NewTaskTag(id TaskTagID, name string, userID UserID) (*TaskTag, error) {
	err := validateTaskTagName(name)
	if err != nil {
		return nil, err
	}

	return &TaskTag{
		id:     id,
		name:   name,
		userID: userID,
	}, nil
}

type TaskTag struct {
	id     TaskTagID
	name   string
	userID UserID
}

type TaskTagRepository interface {
	NextID() TaskTagID
	Store(taskTag *TaskTag) error
	GetByID(id TaskTagID) (*TaskTag, error)
	RemoveByID(id TaskTagID) error
}

func (t *TaskTag) ID() TaskTagID {
	return t.id
}

func (t *TaskTag) Name() string {
	return t.name
}

func (t *TaskTag) UserID() UserID {
	return t.userID
}

func (t *TaskTag) ChangeName(newName string) error {
	err := validateTaskTagName(newName)
	if err != nil {
		return err
	}

	t.name = newName
	return nil
}

func validateTaskTagName(name string) error {
	length := utf8.RuneCountInString(name)
	if length < minTaskTagNameLength || length > maxTaskTagNameLength {
		return ErrTaskTagNameLength
	}
	return nil
}
