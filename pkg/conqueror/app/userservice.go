package app

import (
	stderrors "errors"

	"conqueror/pkg/conqueror/domain"

	"github.com/pkg/errors"
)

var ErrUserAlreadyExists = stderrors.New("user already exists")

type UserService interface {
	RegisterUser(login, password, nickname string) error
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
	user, err := domain.NewUser(userID, login, password, nickname)
	if err != nil {
		return err
	}

	return s.userRepository.Store(user)
}
