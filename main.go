package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/artoriaschan/bloger-server/service"
	"github.com/artoriaschan/bloger-server/utils/logging"
)

var appName = "bloger-server"
var InfoLogger = logging.GetLogger(logging.InfoPath, "Info")

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	InfoLogger.Println("Starting " + appName)
	service.StartWebServer("8088")
}
