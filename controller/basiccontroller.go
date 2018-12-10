package controller

import (
	"encoding/json"

	"github.com/artoriaschan/bloger-server/utils/logging"
	_ "github.com/artoriaschan/bloger-server/utils/memory"
	"github.com/artoriaschan/bloger-server/utils/session"
)

type ResponseResult struct {
	Code    State       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseList struct {
	List       interface{} `json:"list"`
	Pagination Pagination  `json:"pagination"`
}
type Pagination struct {
	Total       int `json:"total"`
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

var AccessLogger = logging.GetLogger(logging.AccessPath, "Info")
var ConsoleLogger = logging.GetConsoleLogger()
var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	var err error
	globalSessions, err = session.NewManager("memory", "GO_SESSION_ID", 3600)
	if err != nil {
		ConsoleLogger.Println(err)
	}
	go globalSessions.GC()
}

// func(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	writer.WriteHeader(http.StatusOK)
// 	writer.Write([]byte("{'path':'/','result':'ok'}"))
// 	//// Case 1: w.Write byte
// 	//w.Write([]byte("Hello World"))
// 	//// Case 2: fmt.Fprintf
// 	//fmt.Fprintf(w, "Hello World")
// 	//// Case 3: io.Write
// 	//io.WriteString(w, "Hello World")
// },

func (rr *ResponseResult) ToJson() []byte {
	resultJson, err := json.Marshal(rr)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
		}
	}()

	return resultJson
}
