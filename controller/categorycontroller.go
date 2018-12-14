package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/artoriaschan/bloger-server/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type ReceiveCategory struct {
	Catename string `json:"catename"`
}

type ReceiveModifyCategory struct {
	Id       string `json:"id"`
	Catename string `json:"catename"`
}

func GetCategories(filter bson.M, writer http.ResponseWriter, request *http.Request) (*[]model.OutputCategory, int, int, int, error) {
	// 获取参数,currentPage和pageSize
	query, _ := url.ParseQuery(request.URL.RawQuery)
	var currentPageQuery string
	var pageSizeQuery string
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
	// 分页处理
	pageSize, _ = strconv.Atoi(pageSizeQuery)
	currentPage, _ = strconv.Atoi(currentPageQuery)
	skip := pageSize * (currentPage - 1)
	limit := pageSize
	cates, num, ok := model.GetCategories(filter, skip, limit)
	if ok {
		return cates, num, currentPage, pageSize, nil
	} else {
		return nil, 0, currentPage, pageSize, nil
	}
}

// /api/cates
func QueryCategories(writer http.ResponseWriter, request *http.Request) {
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
	// 从session中获取登录认证信息
	sess := globalSessions.SessionStart(writer, request)
	isAdmin := sess.Get("isAdmin").(bool) //通过断言实现类型转换
	if isAdmin {
		// 创建筛选条件
		var filter = bson.M{}
		cates, num, currentPage, pageSize, err := GetCategories(filter, writer, request)
		if err != nil {
			panic(err)
		}
		ConsoleLogger.Println(*cates)
		responseResult := ResponseResult{
			Code:    OK,
			Message: "查询成功",
			Data: ResponseList{
				List: *cates,
				Pagination: Pagination{
					Total:       num,
					PageSize:    pageSize,
					CurrentPage: currentPage,
				},
			},
		}
		result := responseResult.ToJson()
		writer.Write(result)
		return
	} else {
		panic(fmt.Errorf("权限不足"))
	}
}

// /api/cate/post
func AddCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(request.Body)
	ConsoleLogger.Println(string(body))

	rc := &ReceiveCategory{}
	err := json.Unmarshal(body, &rc)
	if err != nil {
		panic(err)
	}
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
			return
		}
	}()
	var user model.User
	sess := globalSessions.SessionStart(writer, request)
	if sess.Get("loginAdmin") != nil {
		user = sess.Get("loginAdmin").(model.User)
	} else {
		panic(fmt.Errorf("未登录"))
	}
	category := model.Category{
		Id:         bson.NewObjectId(),
		Catename:   rc.Catename,
		Createtime: time.Now().Unix(),
		Updatetime: time.Now().Unix(),
		Creater:    user.Id,
		Isdelete:   false,
	}
	flag := model.InsertCategory(&category)
	if !flag {
		panic(fmt.Errorf("插入失败"))
	}
	cates, num, currentPage, pageSize, err := GetCategories(bson.M{}, writer, request)
	if err != nil {
		panic(err)
	}
	responseResult := ResponseResult{
		Code:    OK,
		Message: "添加" + rc.Catename + "成功",
		Data: ResponseList{
			List: *cates,
			Pagination: Pagination{
				Total:       num,
				PageSize:    pageSize,
				CurrentPage: currentPage,
			},
		},
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}

// /api/cates/update
func ModifyCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(request.Body)
	ConsoleLogger.Println(string(body))

	rmc := &ReceiveModifyCategory{}
	err := json.Unmarshal(body, &rmc)
	if err != nil {
		panic(err)
	}
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
			return
		}
	}()
	sess := globalSessions.SessionStart(writer, request)
	if sess.Get("loginAdmin") != nil {
		// TODO
	} else {
		panic(fmt.Errorf("未登录"))
	}
	ok := model.UpdateCategory(rmc.Id, rmc.Catename)
	if !ok {
		panic(fmt.Errorf("更新失败"))
	}
	cates, num, currentPage, pageSize, err := GetCategories(bson.M{}, writer, request)
	if err != nil {
		panic(err)
	}
	responseResult := ResponseResult{
		Code:    OK,
		Message: "更新" + rmc.Catename + "成功",
		Data: ResponseList{
			List: *cates,
			Pagination: Pagination{
				Total:       num,
				PageSize:    pageSize,
				CurrentPage: currentPage,
			},
		},
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}

// /api/cate/delete/{cateId}
func DeleteCategory(writer http.ResponseWriter, request *http.Request) {
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
			return
		}
	}()
	vars := mux.Vars(request)
	cateId := vars["cateId"]
	if cateId == "" {
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
		ok := model.DeleteCategory(cateId)
		if !ok {
			panic(fmt.Errorf("更新失败"))
		}
	}
	cates, num, currentPage, pageSize, err := GetCategories(bson.M{}, writer, request)
	if err != nil {
		panic(err)
	}
	responseResult := ResponseResult{
		Code:    OK,
		Message: "删除成功",
		Data: ResponseList{
			List: *cates,
			Pagination: Pagination{
				Total:       num,
				PageSize:    pageSize,
				CurrentPage: currentPage,
			},
		},
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}
