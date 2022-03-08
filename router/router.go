package router

import (
	"api/router/endpoints"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	return endpoints.ConfigRouter(r)
}
