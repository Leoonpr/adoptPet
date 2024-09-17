package routes

import (
	"api/src/controllers"
	"net/http"
)

var adopterRoutes = []Route{
	{
		URI:                    "/adopters",
		Method:                 http.MethodPost,
		Function:               controllers.CreateAdopter,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/adopters",
		Method:                 http.MethodGet,
		Function:               controllers.ReadAdopters,
		RequiresAuthentication: false,
	},
	
}
