package service

import (
	stderrors "errors"
	"io"
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
	ChangeUserAvatar(userID uuid.UUID, file io.Reader) error
	ChangeUserPassword(userID uuid.UUID, newPassword string) error
}

func NewUserService(userRepository domain.UserRepository, imageRepository domain.ImageRepository) UserService {
	return &userService{
		userRepository:  userRepository,
		imageRepository: imageRepository,
	}
}

type userService struct {
	userRepository  domain.UserRepository
	imageRepository domain.ImageRepository
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

	user, err := domain.NewUser(userID, login, passwordHash, nil)
	if err != nil {
		return err
	}

	return s.userRepository.Store(user)
}

func (s *userService) ChangeUserAvatar(userID uuid.UUID, file io.Reader) error {
	existingUser, err := s.userRepository.GetByID(domain.UserID(userID))
	if err != nil {
		return err
	}

	if existingUser.AvatarID() == nil {
		avatarID := s.imageRepository.NextID()
		avatar, err := domain.NewImage(avatarID, uuid.UUID(avatarID).String()+".jpg")
		if err != nil {
			return err
		}

		err = s.imageRepository.Store(avatar, file)
		if err != nil {
			return err
		}

		err = existingUser.ChangeAvatarID(&avatarID)
		if err != nil {
			return err
		}

		return s.userRepository.Store(existingUser)
	}

	existingAvatar, err := s.imageRepository.GetByID(*existingUser.AvatarID())
	if err != nil {
		return err
	}

	return s.imageRepository.Store(existingAvatar, file)
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
