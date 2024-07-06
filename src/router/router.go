package router

import "github.com/gorilla/mux"

func Gerar() *mux.Router {
	router := mux.NewRouter()
	return router
}
