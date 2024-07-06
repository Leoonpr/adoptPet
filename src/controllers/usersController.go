package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating an User"))
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
