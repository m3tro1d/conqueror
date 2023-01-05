package service

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

type NoteService interface {
	CreateNote(userID uuid.UUID, title, content string, subjectID *uuid.UUID) error
	UpdateNote(noteID uuid.UUID, title, content string, subjectID *uuid.UUID) error
	RemoveNote(noteID uuid.UUID) error
}

func NewNoteService(noteRepository domain.NoteRepository, userRepository domain.UserRepository) NoteService {
	return &noteService{
		noteRepository: noteRepository,
		userRepository: userRepository,
	}
}

type noteService struct {
	noteRepository domain.NoteRepository
	userRepository domain.UserRepository
}

func (s *noteService) CreateNote(userID uuid.UUID, title, content string, subjectID *uuid.UUID) error {
	err := validateUserExists(s.userRepository, userID)
	if err != nil {
		return err
	}

	noteID := s.noteRepository.NextID()

	note, err := domain.NewNote(noteID, domain.UserID(userID), title, content, (*domain.SubjectID)(subjectID))
	if err != nil {
		return err
	}

	return s.noteRepository.Store(note)
}

func (s *noteService) UpdateNote(noteID uuid.UUID, title, content string, subjectID *uuid.UUID) error {
	note, err := s.noteRepository.GetByID(domain.NoteID(noteID))
	if err != nil {
		return err
	}

	err = note.ChangeTitle(title)
	if err != nil {
		return err
	}

	err = note.ChangeContent(content)
	if err != nil {
		return err
	}

	err = note.ChangeSubjectID((*domain.SubjectID)(subjectID))
	if err != nil {
		return err
	}

	return s.noteRepository.Store(note)
}

func (s *noteService) RemoveNote(noteID uuid.UUID) error {
	existingNote, err := s.noteRepository.GetByID(domain.NoteID(noteID))
	if err != nil {
		return err
	}

	return s.noteRepository.RemoveByID(existingNote.ID())
}
