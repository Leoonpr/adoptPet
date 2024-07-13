package repositories

import (
	"api/src/models"
	"database/sql"
)

type shelter struct {
	db *sql.DB
}

func NewShelterRepository(db *sql.DB) *shelter {
	return &shelter{db}
}

func (repository shelter) Create(shelter models.Shelter) (uint64, error) {
	statment, err := repository.db.Prepare(
		"INSERT INTO shelter (name, city, address, email, phone, cnpj) VALUES (?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(shelter.Name, shelter.City, shelter.Address, shelter.Email, shelter.Phone, shelter.CNPJ)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}
