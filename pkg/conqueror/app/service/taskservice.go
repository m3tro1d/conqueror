package service

import (
	"time"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
	"github.com/pkg/errors"
)

type TaskService interface {
	CreateTask(userID uuid.UUID, dueDate time.Time, title, description string, subjectID *uuid.UUID) error
	ChangeTaskTitle(taskID uuid.UUID, newTitle string) error
	ChangeTaskDescription(taskID uuid.UUID, newDescription string) error
	ChangeTaskStatus(taskID uuid.UUID, newStatus int) error
	ChangeTaskTags(taskID uuid.UUID, tags []uuid.UUID) error
	RemoveTask(taskID uuid.UUID) error
}

func NewTaskService(taskRepository domain.TaskRepository, userRepository domain.UserRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
		userRepository: userRepository,
	}
}

type taskService struct {
	taskRepository domain.TaskRepository
	userRepository domain.UserRepository
}

func (s *taskService) CreateTask(userID uuid.UUID, dueDate time.Time, title string, description string, subjectID *uuid.UUID) error {
	err := validateUserExists(s.userRepository, userID)
	if err != nil {
		return err
	}

	taskID := s.taskRepository.NextID()

	task, err := domain.NewTask(
		taskID,
		domain.UserID(userID),
		dueDate,
		title,
		description,
		domain.TaskStatusOpen,
		(*domain.SubjectID)(subjectID),
	)
	if err != nil {
		return err
	}

	return s.taskRepository.Store(task)
}

func (s *taskService) ChangeTaskTitle(taskID uuid.UUID, newTitle string) error {
	task, err := s.taskRepository.GetByID(domain.TaskID(taskID))
	if err != nil {
		return err
	}

	err = task.ChangeTitle(newTitle)
	if err != nil {
		return err
	}

	return s.taskRepository.Store(task)
}

func (s *taskService) ChangeTaskDescription(taskID uuid.UUID, newDescription string) error {
	task, err := s.taskRepository.GetByID(domain.TaskID(taskID))
	if err != nil {
		return err
	}

	err = task.ChangeDescription(newDescription)
	if err != nil {
		return err
	}

	return s.taskRepository.Store(task)
}

func (s *taskService) ChangeTaskStatus(taskID uuid.UUID, newStatus int) error {
	task, err := s.taskRepository.GetByID(domain.TaskID(taskID))
	if err != nil {
		return err
	}

	status, err := appToDomainTaskStatus(newStatus)
	if err != nil {
		return err
	}

	err = task.ChangeStatus(status)
	if err != nil {
		return err
	}

	return s.taskRepository.Store(task)
}

func (s *taskService) ChangeTaskTags(taskID uuid.UUID, tags []uuid.UUID) error {
	task, err := s.taskRepository.GetByID(domain.TaskID(taskID))
	if err != nil {
		return err
	}

	err = task.ChangeTags(convertUUIDsToTaskTagIDs(tags))
	if err != nil {
		return err
	}

	return s.taskRepository.Store(task)
}

func (s *taskService) RemoveTask(taskID uuid.UUID) error {
	existingTask, err := s.taskRepository.GetByID(domain.TaskID(taskID))
	if err != nil {
		return err
	}

	return s.taskRepository.RemoveByID(existingTask.ID())
}

func appToDomainTaskStatus(status int) (domain.TaskStatus, error) {
	domainStatus := domain.TaskStatus(status)

	switch domainStatus {
	case domain.TaskStatusOpen,
		domain.TaskStatusCompleted:
		return domainStatus, nil
	default:
		return 0, errors.WithStack(domain.ErrInvalidTaskStatus)
	}
}

func convertUUIDsToTaskTagIDs(tags []uuid.UUID) []domain.TaskTagID {
	result := make([]domain.TaskTagID, 0, len(tags))
	for _, tagID := range tags {
		result = append(result, domain.TaskTagID(tagID))
	}
	return result
}