package domain

import (
	"fmt"
	"unicode/utf8"
)

const (
	minUsernameLength = 1
	maxUsernameLength = 100
)

var ErrUsernameLength = fmt.Errorf("username must be more or equal to %d and less or equal to %d", minUsernameLength, maxUsernameLength)

func NewUser(id UserID, username string, password string) *User {
	return &User{
		id:       id,
		username: username,
		password: password,
	}
}

type User struct {
	id       UserID
	username string
	password string
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) ChangeUsername(newUsername string) error {
	err := validateUsername(newUsername)
	if err != nil {
		return err
	}

	u.username = newUsername
	return nil
}

func validateUsername(username string) error {
	length := utf8.RuneCountInString(username)
	if length < minUsernameLength || length > maxUsernameLength {
		return ErrUsernameLength
	}
	return nil
}
