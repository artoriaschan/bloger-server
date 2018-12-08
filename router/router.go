package router

import (
	"github.com/rs/cors"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {

	router := mux.NewRouter().StrictSlash(true)
	// 设置静态目录
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	for _, route := range routes {
		router.Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	// 添加cors设置
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000","http://localhost:8000"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})
	handler := c.Handler(router)
	return handler
}
