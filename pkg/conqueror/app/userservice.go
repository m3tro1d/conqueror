package app

import "conqueror/pkg/conqueror/domain"

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
	user, err := domain.NewUser(1, login, password, nickname)
	if err != nil {
		return err
	}

	return s.userRepository.Store(user)
}
