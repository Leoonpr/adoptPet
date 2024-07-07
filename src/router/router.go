package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	router := mux.NewRouter()
	return routes.Configurar(router)
}
