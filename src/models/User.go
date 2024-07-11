package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) Prepare(phase string) error {
	if erro := user.validate(phase); erro != nil {
		return erro
	}
	user.format()
	return nil
}

func (user *User) validate(phase string) error {
	if user.Name == "" {
		return errors.New("the name field is required")
	}
	if user.Nick == "" {
		return errors.New("the nick field is required")
	}
	if user.Email == "" {
		return errors.New("the email field is required")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return erro
	}

	if phase == "register" && user.Password == "" {
		return errors.New("the password field is required")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
