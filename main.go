package main

import (
	"github.com/artoriaschan/bloger-server/service"
	"log"
)

var appName = "bloger-server"

func main(){
	log.Printf("Starting %v\n", appName)
	service.StartWebServer("8088")
}
