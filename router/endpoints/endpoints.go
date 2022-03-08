package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Uri          string
	Method       string
	Fun          func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := endpointsUsers

	for _, router := range routes {
		r.HandleFunc(router.Uri, router.Fun).Methods(router.Method)
	}

	return r
}
