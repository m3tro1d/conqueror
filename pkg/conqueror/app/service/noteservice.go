package service

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
)

type NoteService interface {
	CreateNote(userID uuid.UUID, title, content string, subjectID *uuid.UUID) error
	ChangeNoteTitle(noteID uuid.UUID, newTitle string) error
	ChangeNoteContent(noteID uuid.UUID, newContent string) error
	ChangeNoteTags(noteID uuid.UUID, tags []uuid.UUID) error
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

func (s *noteService) ChangeNoteTitle(noteID uuid.UUID, newTitle string) error {
	note, err := s.noteRepository.GetByID(domain.NoteID(noteID))
	if err != nil {
		return err
	}

	err = note.ChangeTitle(newTitle)
	if err != nil {
		return err
	}

	return s.noteRepository.Store(note)
}

func (s *noteService) ChangeNoteContent(noteID uuid.UUID, newContent string) error {
	note, err := s.noteRepository.GetByID(domain.NoteID(noteID))
	if err != nil {
		return err
	}

	err = note.ChangeContent(newContent)
	if err != nil {
		return err
	}

	return s.noteRepository.Store(note)
}

func (s *noteService) ChangeNoteTags(noteID uuid.UUID, tags []uuid.UUID) error {
	note, err := s.noteRepository.GetByID(domain.NoteID(noteID))
	if err != nil {
		return err
	}

	err = note.ChangeTags(convertUUIDsToNoteTagIDs(tags))
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

func convertUUIDsToNoteTagIDs(tags []uuid.UUID) []domain.NoteTagID {
	result := make([]domain.NoteTagID, 0, len(tags))
	for _, tagID := range tags {
		result = append(result, domain.NoteTagID(tagID))
	}
	return result
}
