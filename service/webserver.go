package service

import (
	"github.com/artoriaschan/bloger-server/router"
	"log"
	"net/http"
)

func StartWebServer(port string) {
	//r := router.NewRouter()
	//handler := cors.Default().Handler(mux)
	//http.Handle("/", r)

	handler := router.NewHandler()
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":" + port, handler)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
