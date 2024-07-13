package controllers

import (
	"fmt"
	"net/http"
)

func CreateShelter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a shelter")
}
func ReadShelters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reading all shelters")
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
