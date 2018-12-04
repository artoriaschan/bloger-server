package controller

import (
	"encoding/json"
	"github.com/artoriaschan/bloger-server/utils/logging"
	"log"
)

type ResponseResult struct {
	Code    State       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
var AccessLogger = logging.GetLogger(logging.AccessPath,"Info")
var ConsoleLogger = logging.GetConsoleLogger()
func (rr *ResponseResult) ToJson() []byte {
	resultJson, err := json.Marshal(rr)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	return resultJson
}
