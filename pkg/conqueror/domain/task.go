package domain

import (
	stderrors "errors"
	"fmt"
	"time"
	"unicode/utf8"
)

const (
	minTaskTitleLength       = 1
	maxTaskTitleLength       = 200
	maxTaskDescriptionLength = 1000

	minTaskTagNameLength = 1
	maxTaskTagNameLength = 200
)

var (
	ErrTaskTitleLength       = fmt.Errorf("task title must be greater or equal to %d and less or equal to %d", minTaskTitleLength, maxTaskTitleLength)
	ErrTaskDescriptionLength = fmt.Errorf("task description must be less or equal to %d", maxTaskDescriptionLength)
	ErrTaskTagNameLength     = fmt.Errorf("task tag name must be greater or equal to %d and less or equal to %d", minTaskTagNameLength, maxTaskTagNameLength)

	ErrTaskNotFound = stderrors.New("task not found")
)

func NewTask(id TaskID, userID UserID, dueDate time.Time, title string, description string) (*Task, error) {
	err := validateTaskTitle(title)
	if err != nil {
		return nil, err
	}

	err = validateTaskDescription(description)
	if err != nil {
		return nil, err
	}

	return &Task{
		id:          id,
		userID:      userID,
		dueDate:     dueDate,
		title:       title,
		description: description,
	}, nil
}

type Task struct {
	id          TaskID
	userID      UserID
	dueDate     time.Time
	title       string
	description string
	tags        []TaskTag
}

type TaskTag struct {
	name      string
	subjectID *SubjectID
}

type TaskRepository interface {
	NextID() TaskID
	Store(task *Task) error
	GetByID(id TaskID) (*Task, error)
	RemoveByID(id TaskID) error
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) UserID() UserID {
	return t.userID
}

func (t *Task) DueDate() time.Time {
	return t.dueDate
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Tags() []TaskTag {
	return t.tags
}

func (t *Task) ChangeTitle(newTitle string) error {
	err := validateTaskTitle(newTitle)
	if err != nil {
		return err
	}

	t.title = newTitle
	return nil
}

func (t *Task) ChangeDescription(newDescription string) error {
	err := validateTaskDescription(newDescription)
	if err != nil {
		return err
	}

	t.description = newDescription
	return nil
}

func (t *TaskTag) Name() string {
	return t.name
}

func (t *TaskTag) SubjectID() *SubjectID {
	return t.subjectID
}

func (t *TaskTag) ChangeName(newName string) error {
	err := validateTaskTagName(newName)
	if err != nil {
		return err
	}

	t.name = newName
	return nil
}

func (t *TaskTag) ChangeSubjectID(newSubjectID *SubjectID) error {
	t.subjectID = newSubjectID
	return nil
}

func validateTaskTitle(title string) error {
	length := utf8.RuneCountInString(title)
	if length < minTaskTitleLength || length > maxTaskTitleLength {
		return ErrTaskTitleLength
	}
	return nil
}

func validateTaskDescription(description string) error {
	length := utf8.RuneCountInString(description)
	if length > maxTaskDescriptionLength {
		return ErrTaskDescriptionLength
	}
	return nil
}

func validateTaskTagName(name string) error {
	length := utf8.RuneCountInString(name)
	if length < minTaskTagNameLength || length > maxTaskTagNameLength {
		return ErrTaskTagNameLength
	}
	return nil
}
