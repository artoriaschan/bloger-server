package main

import (
	"fmt"
	"github.com/artoriaschan/bloger-server/service"
	"github.com/artoriaschan/bloger-server/utils/logging"
	"log"
	"os"
	"path/filepath"
)

var appName = "bloger-server"
var InfoLogger = logging.GetLogger(logging.InfoPath,"Info")
func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	InfoLogger.Println("Starting " + appName)
	service.StartWebServer("8088")
}
