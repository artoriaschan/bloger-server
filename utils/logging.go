package utils

import (
	"fmt"
	"os"
	"time"

	logging "github.com/op/go-logging"
)

// 系统访问日志
type AccessLogger struct {
	Type      string `json:"Type"`
	IP        string `json:"IP"`
	Header    string `json:"Header"`
	UserAgent string `json:"UserAgent"`
	Extend    string `json:"Extend"`
}

// 系统正常打印日志
type InfoLogger struct {
	Type    string `json:"Type"`
	Package string `json:"Package"`
	Method  string `json:"Method"`
	Message string `json:"Message"`
}

// 系统异常日志
type ErrorLogger struct {
	Type         string `json:"Type"`
	Package      string `json:"Package"`
	Method       string `json:"Method"`
	ErrorMessage string `json:"ErrorMessage"`
}
type Password string

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func SetLoggerConfig(path, dir string) *logging.Logger {
	// LoggingPath = "/Users/artorias/Desktop/workspace/go/src/github.com/artoriaschan/bloger-server/logs"
	var LoggingPath string = path
	var loggingSubDir string = dir
	var Logger = logging.MustGetLogger("blog-server")

	logFileName := "log_" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(LoggingPath+"/"+loggingSubDir+"/"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.INFO, "") // 设置等级

	logging.SetBackend(backend1Leveled, backend2Formatter)

	return Logger
}

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}
