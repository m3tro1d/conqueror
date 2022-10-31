package app

import (
	"time"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

type TaskService interface {
	CreateTask(userID uuid.UUID, dueDate time.Time, title string, description string, subjectID *uuid.UUID) error
	ChangeTaskTitle(taskID uuid.UUID, newTitle string) error
	ChangeTaskDescription(taskID uuid.UUID, newDescription string) error
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

	task, err := domain.NewTask(taskID, domain.UserID(userID), dueDate, title, description, (*domain.SubjectID)(subjectID))
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

func (s *taskService) RemoveTask(taskID uuid.UUID) error {
	existingTask, err := s.taskRepository.GetByID(domain.TaskID(taskID))
	if err != nil {
		return err
	}

	return s.taskRepository.RemoveByID(existingTask.ID())
}
