package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	var dir string

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	for _, route := range routes {
		router.Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
