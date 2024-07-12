package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var user models.User
	if erro = json.Unmarshal(corpoRequisicao, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	database, erro := db.Conection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer database.Close()

	repositorio := repositories.NewUsersRepository(database)
	usuarioSalvoNoBanco, erro := repositorio.FindByEmail(user.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	fmt.Println("")
	fmt.Println("Hash armazenado no banco:", usuarioSalvoNoBanco.Password)
	fmt.Println("Comprimento do hash:", len(usuarioSalvoNoBanco.Password))
	fmt.Println("Senha fornecida pelo usuário:", user.Password)
	if erro = security.Compare(usuarioSalvoNoBanco.Password, user.Password); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	token, erro := authentication.CreateToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
}
