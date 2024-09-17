package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateAdopter(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	var adopter models.Adopter
	if err = json.Unmarshal(bodyRequest, &adopter); err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err = adopter.Prepare(); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	defer database.Close()
	repository := repositories.NewAdopterRepository(database)
	adopter.ID, err = repository.Create(adopter)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, adopter)

}

func ReadAdopters(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("adopter"))
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	defer database.Close()
	repository := repositories.NewAdopterRepository(database)
	adopters, err := repository.ReadAll(name)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, adopters)
}

func ReadAdopter(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	adopterID, err := strconv.ParseUint(parameters["adopterID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer database.Close()
	repository := repositories.NewAdopterRepository(database)
	adopter, err := repository.ReadByID(adopterID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, adopter)

}

func UpdateAdopter(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	adopterID, err := strconv.ParseUint(parameters["adopterID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var adopter models.Adopter
	if err = json.Unmarshal(bodyRequest, &adopter); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = adopter.Prepare(); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer database.Close()
	repository := repositories.NewAdopterRepository(database)
	if err = repository.Update(adopterID, adopter); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
}

func DeleteAdopter(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	adopterID, err := strconv.ParseUint(parameters["adopterID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	defer database.Close()
	repository := repositories.NewAdopterRepository(database)
	if err = repository.Delete(adopterID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
}
