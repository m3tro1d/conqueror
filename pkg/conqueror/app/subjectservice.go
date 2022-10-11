package app

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

type SubjectService interface {
	CreateSubject(userID uuid.UUID, title string) error
	ChangeSubjectTitle(subjectID uuid.UUID, newTitle string) error
	RemoveSubject(subjectID uuid.UUID) error
}

func NewSubjectService(subjectRepository domain.SubjectRepository, userRepository domain.UserRepository) SubjectService {
	return &subjectService{
		subjectRepository: subjectRepository,
		userRepository:    userRepository,
	}
}

type subjectService struct {
	subjectRepository domain.SubjectRepository
	userRepository    domain.UserRepository
}

func (s *subjectService) CreateSubject(userID uuid.UUID, title string) error {
	err := s.validateUserExists(userID)
	if err != nil {
		return err
	}

	subjectID := s.subjectRepository.NextID()

	subject, err := domain.NewSubject(subjectID, domain.UserID(userID), title)
	if err != nil {
		return err
	}

	return s.subjectRepository.Store(subject)
}

func (s *subjectService) ChangeSubjectTitle(subjectID uuid.UUID, newTitle string) error {
	existingSubject, err := s.subjectRepository.GetByID(domain.SubjectID(subjectID))
	if err != nil {
		return err
	}

	err = existingSubject.ChangeTitle(newTitle)
	if err != nil {
		return err
	}

	return s.subjectRepository.Store(existingSubject)
}

func (s *subjectService) RemoveSubject(subjectID uuid.UUID) error {
	existingSubject, err := s.subjectRepository.GetByID(domain.SubjectID(subjectID))
	if err != nil {
		return err
	}

	return s.subjectRepository.RemoveByID(existingSubject.ID())
}

func (s *subjectService) validateUserExists(userID uuid.UUID) error {
	_, err := s.userRepository.GetByID(domain.UserID(userID))
	return err
}
