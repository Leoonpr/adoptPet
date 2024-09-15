package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

func Configurar(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, shelterRoutes...)
	routes = append(routes, adopterRoutes...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}
	return r
}
