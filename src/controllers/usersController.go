package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}
	database, erro := db.Conection()
	if erro != nil {
		log.Fatal(erro)
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	usuarioID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading an User"))
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
