package repositories

import (
	"api/src/models"
	"database/sql"
)

type adopter struct {
	db *sql.DB
}

func NewAdopterRepository(db *sql.DB) *adopter {
	return &adopter{db}
}

func (repository adopter) Create(newAdopter models.Adopter) (uint64, error) {
	statment, err := repository.db.Prepare(
		"insert into adopter (name, email, cpf, phone) values(?,?,?,?)",
	)

	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(newAdopter.Name, newAdopter.Email, newAdopter.CPF, newAdopter.Phone)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}
