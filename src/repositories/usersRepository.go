package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repositorio users) Create(usuario models.User) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"insert into users (name, nick, email, password) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()
	resultado, erro := statment.Exec(usuario.Name, usuario.Nick, usuario.Email, usuario.Password)
	if erro != nil {
		return 0, erro
	}
	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIdInserido), nil
}
