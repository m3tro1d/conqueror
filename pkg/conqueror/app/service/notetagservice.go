package service

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

type NoteTagService interface {
	CreateNoteTag(userID uuid.UUID, name string) error
	ChangeNoteTagName(noteTagID uuid.UUID, newName string) error
	RemoveNoteTag(noteTagID uuid.UUID) error
}

func NewNoteTagService(noteTagRepository domain.NoteTagRepository, userRepository domain.UserRepository) NoteTagService {
	return &noteTagService{
		noteTagRepository: noteTagRepository,
		userRepository:    userRepository,
	}
}

type noteTagService struct {
	noteTagRepository domain.NoteTagRepository
	userRepository    domain.UserRepository
}

func (s *noteTagService) CreateNoteTag(userID uuid.UUID, name string) error {
	err := validateUserExists(s.userRepository, userID)
	if err != nil {
		return err
	}

	noteTagID := s.noteTagRepository.NextID()

	noteTag, err := domain.NewNoteTag(noteTagID, name, domain.UserID(userID))
	if err != nil {
		return err
	}

	return s.noteTagRepository.Store(noteTag)
}

func (s *noteTagService) ChangeNoteTagName(noteTagID uuid.UUID, newName string) error {
	noteTag, err := s.noteTagRepository.GetByID(domain.NoteTagID(noteTagID))
	if err != nil {
		return err
	}

	err = noteTag.ChangeName(newName)
	if err != nil {
		return err
	}

	return s.noteTagRepository.Store(noteTag)
}

func (s *noteTagService) RemoveNoteTag(noteTagID uuid.UUID) error {
	noteTag, err := s.noteTagRepository.GetByID(domain.NoteTagID(noteTagID))
	if err != nil {
		return err
	}

	return s.noteTagRepository.RemoveByID(noteTag.ID())
}
