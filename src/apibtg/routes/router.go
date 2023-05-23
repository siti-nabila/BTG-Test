package routes

import (
	"github.com/gorilla/mux"
)

func CreateRouter(host, port string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Host(host)
	return router
}
