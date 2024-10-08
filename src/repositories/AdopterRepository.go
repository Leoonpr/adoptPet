package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository adopter) ReadAll(name string) ([]models.Adopter, error) {
	name = fmt.Sprintf("%%%s%%", name)
	lines, err := repository.db.Query("Select id, name, email, cpf, phone from adopter WHERE name LIKE ?", name)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var adopters []models.Adopter
	for lines.Next() {
		var adopter models.Adopter
		if err = lines.Scan(
			&adopter.ID, &adopter.Name, &adopter.Email, &adopter.CPF, &adopter.Phone,
		); err != nil {
			return nil, err
		}
		adopters = append(adopters, adopter)
	}
	return adopters, nil
}
func (repository adopter) ReadByID(id uint64) (models.Adopter, error) {
	lines, err := repository.db.Query("Select id, name, email, cpf, phone from adopter WHERE id =?", id)
	if err != nil {
		return models.Adopter{}, err
	}
	defer lines.Close()
	var adopter models.Adopter
	if lines.Next() {
		if err = lines.Scan(
			&adopter.ID, &adopter.Name, &adopter.Email, &adopter.CPF, &adopter.Phone,
		); err != nil {
			return models.Adopter{}, err
		}
	}
	return adopter, nil
}

func (repository adopter) Update(id uint64, newAdopter models.Adopter) error {
	statment, err := repository.db.Prepare(
		"UPDATE adopter SET name =?, email =?, cpf =?, phone =? WHERE id =?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()
	if _, err = statment.Exec(newAdopter.Name, newAdopter.Email, newAdopter.CPF, newAdopter.Phone, id); err != nil {
		return err
	}
	return nil
}

func (repository adopter) Delete(id uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM adopter WHERE id =?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}
	return nil
}
