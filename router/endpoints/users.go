package endpoints

import (
	"api/controllers"
	"net/http"
)

var endpointsUsers = []Router{
	{
		Uri:          "/users",
		Method:       http.MethodPost,
		Fun:          controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		Uri:          "/users",
		Method:       http.MethodGet,
		Fun:          controllers.GetUsers,
		RequiredAuth: false,
	},
	{
		Uri:          "/users/{id}",
		Method:       http.MethodGet,
		Fun:          controllers.GetOneUser,
		RequiredAuth: false,
	},
	{
		Uri:          "/users/{id}",
		Method:       http.MethodPut,
		Fun:          controllers.PutUser,
		RequiredAuth: false,
	},
	{
		Uri:          "/users/{id}",
		Method:       http.MethodDelete,
		Fun:          controllers.DeleteUser,
		RequiredAuth: false,
	},
}
