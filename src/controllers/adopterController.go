package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
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
