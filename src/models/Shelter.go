package models

import (
	"errors"
	"strings"
)

type Shelter struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	City      string `json:"city,omitempty"`
	Address   string `json:"address,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CNPJ      string `json:"cnpj,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (shelter *Shelter) Prepare() error {
	if err := shelter.validate(); err != nil {
		return err
	}
	if err := shelter.format(); err != nil {
		return err
	}
	return nil
}

func (shelter *Shelter) validate() error {
	if shelter.Name == "" {
		return errors.New("the name field is required")
	}
	if shelter.City == "" {
		return errors.New("the city field is required")
	}
	if shelter.Address == "" {
		return errors.New("the address field is required")
	}
	if shelter.Email == "" {
		return errors.New("the email field is required")
	}
	if shelter.Phone == "" {
		return errors.New("the phone field is required")
	}
	if shelter.CNPJ == "" {
		return errors.New("the cnpj field is required")
	}
	return nil
}

func (shelter *Shelter) format() error {
	shelter.Name = strings.TrimSpace(shelter.Name)
	shelter.City = strings.TrimSpace(shelter.City)
	shelter.Address = strings.TrimSpace(shelter.Address)
	shelter.Email = strings.TrimSpace(shelter.Email)
	shelter.Phone = strings.TrimSpace(shelter.Phone)
	shelter.CNPJ = strings.TrimSpace(shelter.CNPJ)
	return nil
}
