package models

import (
	"api/src/security"
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
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(phase string) error {
	if err := user.validate(phase); err != nil {
		return err
	}
	if err := user.format(phase); err != nil {
		return err
	}
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

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return err
	}

	if phase == "register" && user.Password == "" {
		return errors.New("the password field is required")
	}
	return nil
}

func (user *User) format(phase string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if phase == "register" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}
	return nil
}
