package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository users) ReadAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()
	var users []models.User
	for lines.Next() {
		var user models.User
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}
