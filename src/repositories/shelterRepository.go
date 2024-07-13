package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
func (repository shelter) ReadAll(name string) ([]models.Shelter, error) {
	name = fmt.Sprintf("%%%s%%", name)
	lines, err := repository.db.Query("SELECT id, name, city, address, email, phone, cnpj, createdAt FROM shelter WHERE name LIKE ?", name)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var shelters []models.Shelter
	for lines.Next() {
		var shelter models.Shelter
		if err = lines.Scan(
			&shelter.ID, &shelter.Name, &shelter.City, &shelter.Address, &shelter.Email, &shelter.Phone, &shelter.CNPJ, &shelter.CreatedAt,
		); err != nil {
			return nil, err
		}
		shelters = append(shelters, shelter)
	}
	return shelters, nil
}
func (repository shelter) ReadShelterByID(shelterID uint64) (models.Shelter, error) {
	lines, err := repository.db.Query("SELECT id, name, city, address, email, phone, cnpj, createdAt FROM shelter WHERE ID = ?", shelterID)
	if err != nil {
		return models.Shelter{}, err
	}
	defer lines.Close()
	var shelter models.Shelter
	if lines.Next() {
		if err = lines.Scan(
			&shelter.ID, &shelter.Name, &shelter.City, &shelter.Address, &shelter.Email, &shelter.Phone, &shelter.CNPJ, &shelter.CreatedAt,
		); err != nil {
			return models.Shelter{}, err
		}
	}
	return shelter, nil
}
