package domain

import (
	stderrors "errors"
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

const (
	minLoginLength = 1
	maxLoginLength = 50

	minNicknameLength = 1
	maxNicknameLength = 50
)

var ErrLoginLength = fmt.Errorf("login must be more or equal to %d and less or equal to %d", minLoginLength, maxLoginLength)
var ErrNicknameLength = fmt.Errorf("nickname must be more or equal to %d and less or equal to %d", minNicknameLength, maxNicknameLength)
var ErrUserNotFound = stderrors.New("user not found")

func NewUser(id UserID, login, password, nickname string) (*User, error) {
	err := validateLogin(login)
	if err != nil {
		return nil, err
	}

	err = validateNickname(nickname)
	if err != nil {
		return nil, err
	}

	return &User{
		id:       id,
		login:    login,
		password: password,
		nickname: nickname,
	}, nil
}

type User struct {
	id       UserID
	login    string
	password string
	nickname string
}

type UserRepository interface {
	NextID() UserID
	Store(user *User) error
	GetById(id UserID) (*User, error)
	FindByLogin(login string) (*User, error)
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Login() string {
	return u.login
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Nickname() string {
	return u.nickname
}

func (u *User) ChangeLogin(newUsername string) error {
	err := validateLogin(newUsername)
	if err != nil {
		return err
	}

	u.login = newUsername
	return nil
}

func (u *User) ChangePassword(newPassword string) error {
	u.password = newPassword
	return nil
}

func (u *User) ChangeNickname(newNickname string) error {
	err := validateNickname(newNickname)
	if err != nil {
		return err
	}

	u.login = newNickname
	return nil
}

func validateLogin(login string) error {
	length := utf8.RuneCountInString(login)
	if length < minLoginLength || length > maxLoginLength {
		return errors.WithStack(ErrLoginLength)
	}
	return nil
}

func validateNickname(nickname string) error {
	length := utf8.RuneCountInString(nickname)
	if length < minNicknameLength || length > maxNicknameLength {
		return errors.WithStack(ErrNicknameLength)
	}
	return nil
}
