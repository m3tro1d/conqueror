package service

import (
	stderrors "errors"
	"unicode/utf8"

	"conqueror/pkg/common/md5"
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/pkg/errors"
)

const (
	minPasswordLength = 6
)

var ErrUserAlreadyExists = stderrors.New("user already exists")
var ErrWeakPassword = errors.Errorf("password must be greater or equal to %d", minPasswordLength)

type UserService interface {
	RegisterUser(login, password string) error
	ChangeUserPassword(userID uuid.UUID, newPassword string) error
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct {
	userRepository domain.UserRepository
}

func (s *userService) RegisterUser(login, password string) error {
	err := s.validateUserDoesNotExist(login)
	if err != nil {
		return err
	}

	err = validatePassword(password)
	if err != nil {
		return err
	}

	userID := s.userRepository.NextID()
	passwordHash := md5.Hash(password)

	user, err := domain.NewUser(userID, login, passwordHash)
	if err != nil {
		return err
	}

	return s.userRepository.Store(user)
}

func (s *userService) ChangeUserPassword(userID uuid.UUID, newPassword string) error {
	existingUser, err := s.userRepository.GetByID(domain.UserID(userID))
	if err != nil {
		return err
	}

	err = existingUser.ChangePassword(md5.Hash(newPassword))
	if err != nil {
		return err
	}

	return s.userRepository.Store(existingUser)
}

func (s *userService) validateUserDoesNotExist(login string) error {
	existingUser, err := s.userRepository.FindByLogin(login)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.WithStack(ErrUserAlreadyExists)
	}
	return nil
}

func validatePassword(password string) error {
	length := utf8.RuneCountInString(password)
	if length < minPasswordLength {
		return errors.WithStack(ErrWeakPassword)
	}
	return nil
}
