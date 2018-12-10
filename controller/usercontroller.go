package controller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/artoriaschan/bloger-server/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type UserWithToken struct {
	model.User
	Authorization string `json:"authorization"`
}

var pageSize = 10
var currentPage = 1

func CurrentAdmin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
			responseResult := ResponseResult{
				Code:    BadServer,
				Message: "服务异常,请稍后再试",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
		}
	}()
	sess := globalSessions.SessionStart(writer, request)
	// TODO 判断是否为空
	admin := sess.Get("loginAdmin").(model.User)
	isAdmin := sess.Get("isAdmin").(bool) //通过断言实现类型转换
	if !isAdmin {
		responseResult := ResponseResult{
			Code:    NoRight,
			Message: "该账号无权限,请更换账号进行操作",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
		return
	}
	ConsoleLogger.Println("admin", admin)
	// user, ok := model.GetUserById(adminId)
	responseResult := ResponseResult{
		Code:    OK,
		Message: "查询成功",
		Data:    admin.ToOutputUser(),
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}
func CurrentUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
			responseResult := ResponseResult{
				Code:    BadServer,
				Message: "服务异常,请稍后再试",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
		}
	}()
	sess := globalSessions.SessionStart(writer, request)
	// TODO 判断是否为空
	if sess.Get("loginUser") == nil {
		responseResult := ResponseResult{
			Code:    NoLogin,
			Message: "您未登录,请登录账号",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
		return
	}
	user := sess.Get("loginUser").(model.User)
	// user, ok := model.GetUserById(adminId)
	responseResult := ResponseResult{
		Code:    OK,
		Message: "查询成功",
		Data:    user.ToOutputUser(),
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}
func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
			responseResult := ResponseResult{
				Code:    BadServer,
				Message: "服务异常,请稍后再试",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
		}
	}()
	// 获取参数,currentPage和pageSize
	query, _ := url.ParseQuery(request.URL.RawQuery)
	var currentPageQuery string
	var pageSizeQuery string
	var accountTypeQuery string
	var freezenQuery string
	if len(query["currentPage"]) != 0 {
		currentPageQuery = query["currentPage"][0]
	} else {
		currentPageQuery = "1"
	}
	if len(query["pageSize"]) != 0 {
		pageSizeQuery = query["pageSize"][0]
	} else {
		pageSizeQuery = "10"
	}
	if len(query["type"]) != 0 {
		accountTypeQuery = query["type"][0]
	}
	if len(query["freezen"]) != 0 {
		freezenQuery = query["freezen"][0]
	}
	// 从session中获取登录认证信息
	sess := globalSessions.SessionStart(writer, request)
	isAdmin := sess.Get("isAdmin").(bool) //通过断言实现类型转换
	if isAdmin {
		// 分页处理
		pageSize, _ = strconv.Atoi(pageSizeQuery)
		currentPage, _ = strconv.Atoi(currentPageQuery)
		skip := pageSize * (currentPage - 1)
		limit := pageSize
		// 创建筛选条件
		var filter = bson.M{}
		if accountTypeQuery != "" {
			typeInt, _ := strconv.Atoi(accountTypeQuery)
			filter["type"] = typeInt
		}
		if freezenQuery != "" {
			freezenBool, _ := strconv.ParseBool(freezenQuery)
			filter["freezen"] = freezenBool
		}
		ConsoleLogger.Println(filter, skip, limit)
		users, ok := model.GetUsers(filter, skip, limit)
		if ok {
			responseResult := ResponseResult{
				Code:    OK,
				Message: "查询成功",
				Data: ResponseList{
					List: *users,
					Pagination: Pagination{
						Total:       len(*users),
						PageSize:    pageSize,
						CurrentPage: currentPage,
					},
				},
			}
			result := responseResult.ToJson()
			writer.Write(result)
			return
		} else {
			responseResult := ResponseResult{
				Code:    BadDB,
				Message: "查询失败",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
			return
		}
	} else {
		responseResult := ResponseResult{
			Code:    NoRight,
			Message: "该账号无权限,请更换账号进行操作",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
	}
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
			responseResult := ResponseResult{
				Code:    BadServer,
				Message: "服务异常,请稍后再试",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
		}
	}()
	vars := mux.Vars(request)
	userId := vars["userId"]
	if userId == "" {
		writer.WriteHeader(http.StatusNotFound)
		responseResult := ResponseResult{
			Code:    BadServer,
			Message: "未知路径",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
	}
	// 从session中获取登录认证信息
	sess := globalSessions.SessionStart(writer, request)
	isAdmin := sess.Get("isAdmin").(bool) //通过断言实现类型转换
	if isAdmin {
		flag := model.DeleteUser(userId)
		if flag {
			skip := pageSize * (currentPage - 1)
			limit := pageSize
			users, ok := model.GetUsers(bson.M{}, skip, limit)
			if ok {
				responseResult := ResponseResult{
					Code:    OK,
					Message: "删除成功",
					Data: ResponseList{
						List: *users,
						Pagination: Pagination{
							Total:       len(*users),
							PageSize:    pageSize,
							CurrentPage: currentPage,
						},
					},
				}
				result := responseResult.ToJson()
				writer.Write(result)
				return
			} else {
				responseResult := ResponseResult{
					Code:    BadDB,
					Message: "查询失败",
					Data:    nil,
				}
				result := responseResult.ToJson()
				writer.Write(result)
				return
			}
		} else {
			responseResult := ResponseResult{
				Code:    BadDB,
				Message: "删除失败",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
			return
		}
	} else {
		responseResult := ResponseResult{
			Code:    NoRight,
			Message: "该账号无权限,请更换账号进行操作",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
	}
}
func FreezeUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer func() {
		if err := recover(); err != nil {
			ConsoleLogger.Println(err)
			responseResult := ResponseResult{
				Code:    BadServer,
				Message: "服务异常,请稍后再试",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
		}
	}()
	vars := mux.Vars(request)
	userId := vars["userId"]
	if userId == "" {
		writer.WriteHeader(http.StatusNotFound)
		responseResult := ResponseResult{
			Code:    BadServer,
			Message: "未知路径",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
	}
	// 从session中获取登录认证信息
	sess := globalSessions.SessionStart(writer, request)
	isAdmin := sess.Get("isAdmin").(bool) //通过断言实现类型转换
	if isAdmin {
		flag := model.FreezeUser(userId)
		if flag {
			skip := pageSize * (currentPage - 1)
			limit := pageSize
			users, ok := model.GetUsers(bson.M{}, skip, limit)
			if ok {
				responseResult := ResponseResult{
					Code:    OK,
					Message: "冻结成功",
					Data: ResponseList{
						List: *users,
						Pagination: Pagination{
							Total:       len(*users),
							PageSize:    pageSize,
							CurrentPage: currentPage,
						},
					},
				}
				result := responseResult.ToJson()
				writer.Write(result)
				return
			} else {
				responseResult := ResponseResult{
					Code:    BadDB,
					Message: "查询失败",
					Data:    nil,
				}
				result := responseResult.ToJson()
				writer.Write(result)
				return
			}
		} else {
			responseResult := ResponseResult{
				Code:    BadDB,
				Message: "冻结失败",
				Data:    nil,
			}
			result := responseResult.ToJson()
			writer.Write(result)
			return
		}
	} else {
		responseResult := ResponseResult{
			Code:    NoRight,
			Message: "该账号无权限,请更换账号进行操作",
			Data:    nil,
		}
		result := responseResult.ToJson()
		writer.Write(result)
	}
}
