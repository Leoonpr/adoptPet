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

func (repository users) ReadUserByID(userID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		`SELECT id, name, nick, email, createdAt FROM users WHERE id = ?`, userID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()
	var user models.User
	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}

func (repository users) UpdateUser(userID uint64, user models.User) error {
	statment, erro := repository.db.Prepare(
		"UPDATE users SET name =?, nick =?, email =? WHERE id =?",
	)
	if erro != nil {
		return erro
	}
	defer statment.Close()
	if _, erro = statment.Exec(user.Name, user.Nick, user.Email, userID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio users) Delete(userID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"DELETE FROM users WHERE id =?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}
	return nil
}

func (repository users) FindByEmail(email string) (models.User, error) {
	lines, erro := repository.db.Query("SELECT id, password FROM users WHERE email= ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()
	var user models.User
	if lines.Next() {
		if erro = lines.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}
