package models

import (
	"errors"
	"strings"
)

type Adopter struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	CPF       string `json:"cpf,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (adopter *Adopter) Prepare() error {
	if err := adopter.validate(); err != nil {
		return err
	}
	if err := adopter.format(); err != nil {
		return err
	}
	return nil
}

func (adopter *Adopter) validate() error {
	if adopter.Name == "" {
		return errors.New("the name field is required")
	}
	if adopter.Email == "" {
		return errors.New("the email field is required")
	}
	if adopter.CPF == "" {
		return errors.New("the cpf field is required")
	}
	if adopter.Phone == "" {
		return errors.New("the phone field is required")
	}
	return nil
}

func (adopter *Adopter) format() error {
	adopter.Name = strings.TrimSpace(adopter.Name)
	adopter.Email = strings.TrimSpace(adopter.Email)
	adopter.CPF = strings.TrimSpace(adopter.CPF)
	adopter.Phone = strings.TrimSpace(adopter.Phone)
	return nil
}
