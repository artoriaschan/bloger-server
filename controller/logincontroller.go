package controller

import (
	"fmt"
	"github.com/artoriaschan/bloger-server/model"
	"log"
	"net/http"
	"net/url"
)
func userLogin(username, password string) ([]byte){
	var result []byte
	user := model.User{}
	hasUser := model.GetUserByUsername(username, &user)
	if(!hasUser) {
		responseResult := ResponseResult{
			Code: NoRegister,
			Message: "该用户没有注册",
			Data: nil,
		}
		result = responseResult.ToJson()
		return result
	}
	flag := user.CheckPassword(password)
	if(flag) {
		responseResult := ResponseResult{
			Code: OK,
			Message: "查询成功",
			Data: user,
		}
		result = responseResult.ToJson()
	}else{
		responseResult := ResponseResult{
			Code: WrongPassword,
			Message: "用户名/密码输入错误",
			Data: nil,
		}
		result = responseResult.ToJson()
	}
	return result
}
func Login(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	var username string
	var password string
	if request.Method == http.MethodGet {
		query := request.URL.RawQuery
		queryMap, _ := url.ParseQuery(query)
		fmt.Println(queryMap)
		username = queryMap["username"][0]
		password = queryMap["password"][0]
	}
	if request.Method == http.MethodPost {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		username = request.Form.Get("username")
		password = request.Form.Get("password")
	}

	writer.Write(userLogin(username, password))
}
