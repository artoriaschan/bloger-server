package main

import (
	"github.com/artoriaschan/bloger-server/service"
	"github.com/artoriaschan/bloger-server/utils/logging"
)

var appName = "bloger-server"
var InfoLogger = logging.GetLogger(logging.InfoPath,"Info")
func main() {
	InfoLogger.Println("Starting " + appName)
	service.StartWebServer("8088")
}
