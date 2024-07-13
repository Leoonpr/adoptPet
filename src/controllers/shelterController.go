package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateShelter(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	var shelter models.Shelter
	if err = json.Unmarshal(bodyRequest, &shelter); err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err := shelter.Prepare(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	defer database.Close()
	repository := repositories.NewShelterRepository(database)
	shelter.ID, err = repository.Create(shelter)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, shelter)
}

func ReadShelters(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(r.URL.Query().Get("shelter"))
	database, err := db.Conection()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	defer database.Close()
	repository := repositories.NewShelterRepository(database)
	shelters, err := repository.ReadAll(name)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, shelters)

}
func ReadShelter(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	shelterID, err := strconv.ParseUint(parameters["shelterID"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	database, erro := db.Conection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()
	repository := repositories.NewShelterRepository(database)
	shelter, err := repository.ReadShelterByID(shelterID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, shelter)

}
func UpdateShelter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a shelter")
}
func DeleteShelter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a shelter")
}
