package routes

import (
	"api/src/controllers"
	"net/http"
)

var shelterRoutes = []Route{
	{
		URI:                    "/shelters",
		Method:                 http.MethodPost,
		Function:               controllers.CreateShelter,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/shelters/",
		Method:                 http.MethodGet,
		Function:               controllers.ReadShelters,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/shelters{shelterID}",
		Method:                 http.MethodGet,
		Function:               controllers.ReadShelter,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/shelters/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateShelter,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/shelters/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteShelter,
		RequiresAuthentication: true,
	},
}
