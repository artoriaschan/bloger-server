package controller

import (
	"encoding/json"

	"github.com/artoriaschan/bloger-server/utils/jwt"
	"github.com/artoriaschan/bloger-server/utils/logging"
)

type ResponseResult struct {
	Code    State       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var AccessLogger = logging.GetLogger(logging.AccessPath, "Info")
var ConsoleLogger = logging.GetConsoleLogger()

// jwt 设置
var JWTAlg = "HS256"
var JWTTyp = "JWT"

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
func JWTCreator(payload jwtoken.PayLoad) string {
	jwt := jwtoken.JWT{}
	jwt.Header = jwtoken.Header{
		Alg: JWTAlg,
		Typ: JWTTyp,
	}
	jwt.PayLoad = payload
	JWToken := jwt.Encode()
	return JWToken
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
