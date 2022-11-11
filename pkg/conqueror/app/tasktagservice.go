package app

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

type TaskTagService interface {
	CreateTaskTag(userID uuid.UUID, name string) error
	ChangeTaskTagName(taskTagID uuid.UUID, newName string) error
	RemoveTaskTag(taskTagID uuid.UUID) error
}

func NewTaskTagService(taskTagRepository domain.TaskTagRepository, userRepository domain.UserRepository) TaskTagService {
	return &taskTagService{
		taskTagRepository: taskTagRepository,
		userRepository:    userRepository,
	}
}

type taskTagService struct {
	taskTagRepository domain.TaskTagRepository
	userRepository    domain.UserRepository
}

func (s *taskTagService) CreateTaskTag(userID uuid.UUID, name string) error {
	err := validateUserExists(s.userRepository, userID)
	if err != nil {
		return err
	}

	taskTagID := s.taskTagRepository.NextID()

	taskTag, err := domain.NewTaskTag(taskTagID, name, domain.UserID(userID))
	if err != nil {
		return err
	}

	return s.taskTagRepository.Store(taskTag)
}

func (s *taskTagService) ChangeTaskTagName(taskTagID uuid.UUID, newName string) error {
	taskTag, err := s.taskTagRepository.GetByID(domain.TaskTagID(taskTagID))
	if err != nil {
		return err
	}

	err = taskTag.ChangeName(newName)
	if err != nil {
		return err
	}

	return s.taskTagRepository.Store(taskTag)
}

func (s *taskTagService) RemoveTaskTag(taskTagID uuid.UUID) error {
	taskTag, err := s.taskTagRepository.GetByID(domain.TaskTagID(taskTagID))
	if err != nil {
		return err
	}

	return s.taskTagRepository.RemoveByID(taskTag.ID())
}
