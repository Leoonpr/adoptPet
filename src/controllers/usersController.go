package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	if erro = user.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	database, erro := db.Conection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	database, erro := db.Conection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()
	repository := repositories.NewUsersRepository(database)
	users, erro := repository.ReadAll(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, users)

}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading all Users"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a User"))
}
