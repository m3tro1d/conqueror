package app

import (
	"crypto/md5"
	"encoding/hex"
	stderrors "errors"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"

	"github.com/pkg/errors"
)

var ErrUserAlreadyExists = stderrors.New("user already exists")

type UserService interface {
	RegisterUser(login, password, nickname string) error
	ChangeUserPassword(userID uuid.UUID, newPassword string) error
	ChangeUserNickname(userID uuid.UUID, newNickname string) error
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct {
	userRepository domain.UserRepository
}

func (s *userService) RegisterUser(login, password, nickname string) error {
	existingUser, err := s.userRepository.FindByLogin(login)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.WithStack(ErrUserAlreadyExists)
	}

	userID := s.userRepository.NextID()
	passwordHash := hashPassword(password)

	user, err := domain.NewUser(userID, login, passwordHash, nickname)
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

	err = existingUser.ChangePassword(hashPassword(newPassword))
	if err != nil {
		return err
	}

	return s.userRepository.Store(existingUser)
}

func (s *userService) ChangeUserNickname(userID uuid.UUID, newNickname string) error {
	existingUser, err := s.userRepository.GetByID(domain.UserID(userID))
	if err != nil {
		return err
	}

	err = existingUser.ChangeNickname(newNickname)
	if err != nil {
		return err
	}

	return s.userRepository.Store(existingUser)
}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
