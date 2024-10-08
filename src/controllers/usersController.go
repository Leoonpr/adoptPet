package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	if erro = user.Prepare("register"); erro != nil {
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
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
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
	user, erro := repository.ReadUserByID(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("it is not possible to update a user other than yourself"))
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = user.Prepare("update"); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	database, erro := db.Conection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()
	repository := repositories.NewUsersRepository(database)
	if erro = repository.UpdateUser(userID, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("it is not possible to delete a user other than yourself"))
		return
	}

	database, erro := db.Conection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
	}
	defer database.Close()
	repository := repositories.NewUsersRepository(database)
	if erro = repository.Delete(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if userIDToken != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("it is not possible to update a password of a user other than yourself"))
		return
	}
	bodyRequest, err := io.ReadAll(r.Body)

	var password models.Password
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	passwordSavedDB, err := repository.FindPassword(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.Compare(passwordSavedDB, password.Current); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	passwordHash, err := security.Hash(password.New)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repository.UpdatePassword(userID, string(passwordHash)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
