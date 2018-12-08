package controller

import (
	"net/http"

	"github.com/artoriaschan/bloger-server/model"
	"github.com/artoriaschan/bloger-server/utils/jwt"
)

func CurrentAdmin(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
		}
	}()
	goJWT, err := request.Cookie("go_jwt")
	ConsoleLogger.Println(goJWT)
	if err != nil {
		panic(err)
	}
	jwt := jwtoken.JWT{}
	ok := jwt.Decode(goJWT.Value)
	if ok {
		user, ok := model.GetUserById(jwt.Aud)
		ConsoleLogger.Println(jwt.Aud, user, ok)
		if ok {
			responseResult := ResponseResult{
				Code:    OK,
				Message: "查询成功",
				Data:    user,
			}
			result := responseResult.ToJson()
			writer.Write(result)
			return
		}
		responseResult := ResponseResult{
			Code:    BadDB,
			Message: "查询失败",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
		return
	} else {
		responseResult := ResponseResult{
			Code:    ExpireToken,
			Message: "登录信息已失效,请重新登录",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
		return
	}
}
