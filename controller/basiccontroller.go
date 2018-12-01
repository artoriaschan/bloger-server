package controller

import (
	"encoding/json"
	"log"

	"github.com/artoriaschan/bloger-server/utils"
)

var Logger = utils.Logger

type ResponseResult struct {
	Code    State       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
