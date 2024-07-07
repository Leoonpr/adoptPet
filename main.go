package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Printf("Escutando na porta %d", config.Porta)
	portaString := ":" + strconv.Itoa(config.Porta)
	log.Fatal(http.ListenAndServe(portaString, r))

}
