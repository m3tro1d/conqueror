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
)

var ErrLoginLength = fmt.Errorf("login must be greater or equal to %d and less or equal to %d", minLoginLength, maxLoginLength)
var ErrUserNotFound = stderrors.New("user not found")

func NewUser(id UserID, login, password string) (*User, error) {
	err := validateLogin(login)
	if err != nil {
		return nil, err
	}

	return &User{
		id:       id,
		login:    login,
		password: password,
	}, nil
}

type User struct {
	id       UserID
	login    string
	password string
}

type UserRepository interface {
	NextID() UserID
	Store(user *User) error
	GetByID(id UserID) (*User, error)
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

func (u *User) ChangePassword(newPassword string) error {
	u.password = newPassword
	return nil
}

func validateLogin(login string) error {
	length := utf8.RuneCountInString(login)
	if length < minLoginLength || length > maxLoginLength {
		return errors.WithStack(ErrLoginLength)
	}
	return nil
}
