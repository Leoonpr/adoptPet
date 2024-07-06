package routes

import (
	"api/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 "http.MethodPost",
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 "http.MethodGet",
		Function:               controllers.ReadUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 "http.MethodGet",
		Function:               controllers.ReadUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 "http.MethodPut",
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 "http.MethodDelete",
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
