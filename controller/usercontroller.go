package controller

import (
	"net/http"

	"github.com/artoriaschan/bloger-server/model"
)

type UserWithToken struct {
	model.User
	Authorization string `json:"authorization"`
}

func CurrentAdmin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
			responseResult := ResponseResult{
				Code:    NoToken,
				Message: "登录信息已失效,请重新登录",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
		}
	}()
	sess := globalSessions.SessionStart(writer, request)
	adminId := sess.Get("loginAdmin")
	ConsoleLogger.Println("adminId", adminId)
	user, ok := model.GetUserById(adminId)
	if ok {
		responseResult := ResponseResult{
			Code:    OK,
			Message: "查询成功",
			Data:    *user,
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
}
