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
	"strings"
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
	fmt.Println("Reading a shelter")
}
func UpdateShelter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a shelter")
}
func DeleteShelter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a shelter")
}
