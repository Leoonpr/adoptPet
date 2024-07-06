package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API rodando na porta 3333")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3333", r))
}
