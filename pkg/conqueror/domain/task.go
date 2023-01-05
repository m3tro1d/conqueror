package domain

import (
	stderrors "errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	minTaskTitleLength       = 1
	maxTaskTitleLength       = 200
	maxTaskDescriptionLength = 1000
)

var (
	ErrTaskTitleLength       = fmt.Errorf("task title must be greater or equal to %d and less or equal to %d", minTaskTitleLength, maxTaskTitleLength)
	ErrTaskDescriptionLength = fmt.Errorf("task description must be less or equal to %d", maxTaskDescriptionLength)
	ErrDuplicateTaskTags     = stderrors.New("duplicate task tags")
	ErrTaskNotFound          = stderrors.New("task not found")
	ErrInvalidTaskStatus     = stderrors.New("invalid task status")
)

func NewTask(
	id TaskID,
	userID UserID,
	dueDate time.Time,
	title string,
	description string,
	status TaskStatus,
	subjectID *SubjectID,
) (*Task, error) {
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
		status:      status,
		tags:        nil,
		subjectID:   subjectID,
	}, nil
}

type Task struct {
	id          TaskID
	userID      UserID
	dueDate     time.Time
	title       string
	description string
	status      TaskStatus
	tags        []TaskTagID
	subjectID   *SubjectID
}

type TaskStatus int

const (
	TaskStatusOpen = TaskStatus(iota)
	TaskStatusCompleted
)

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

func (t *Task) Status() TaskStatus {
	return t.status
}

func (t *Task) Tags() []TaskTagID {
	return t.tags
}

func (t *Task) SubjectID() *SubjectID {
	return t.subjectID
}

func (t *Task) ChangeDueDate(newDueDate time.Time) error {
	t.dueDate = newDueDate
	return nil
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

func (t *Task) ChangeSubjectID(subjectID *SubjectID) error {
	t.subjectID = subjectID
	return nil
}

func (t *Task) ChangeStatus(newStatus TaskStatus) error {
	t.status = newStatus
	return nil
}

func (t *Task) ChangeTags(tags []TaskTagID) error {
	err := validateTaskTags(tags)
	if err != nil {
		return err
	}

	t.tags = tags
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

func validateTaskTags(tags []TaskTagID) error {
	tagsMap := make(map[TaskTagID]bool)
	for _, tagID := range tags {
		if tagsMap[tagID] {
			return errors.WithStack(ErrDuplicateTaskTags)
		}
		tagsMap[tagID] = true
	}
	return nil
}
