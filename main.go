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
	fmt.Println("API rodando")
	portaString := ":" + strconv.Itoa(config.Porta)
	log.Fatal(http.ListenAndServe(portaString, r))

}
