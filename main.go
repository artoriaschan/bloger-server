package main

import (
	"github.com/artoriaschan/bloger-server/service"
	"github.com/artoriaschan/bloger-server/utils"
)

var appName = "bloger-server"
var Logger = utils.Logger

func main() {
	Logger.Info("Starting " + appName)
	service.StartWebServer("8088")
}
