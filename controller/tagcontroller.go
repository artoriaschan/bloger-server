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
	"gopkg.in/mgo.v2/bson"
)

type ReceiveTag struct {
	Tagname string `json:"tagname"`
	Color   string `json:"color"`
}

func GetTags(filter bson.M, writer http.ResponseWriter, request *http.Request) (*[]model.OutputTag, int, int, int, error) {
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
	tags, num, ok := model.GetTags(filter, skip, limit)
	if ok {
		return tags, num, currentPage, pageSize, nil
	} else {
		return nil, 0, currentPage, pageSize, nil
	}
}

// /api/tags
func QueryTags(writer http.ResponseWriter, request *http.Request) {
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
		tags, num, currentPage, pageSize, err := GetTags(filter, writer, request)
		if err != nil {
			panic(err)
		}
		ConsoleLogger.Println(*tags)
		responseResult := ResponseResult{
			Code:    OK,
			Message: "查询成功",
			Data: ResponseList{
				List: *tags,
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

// /api/tag/post
func AddTag(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(request.Body)
	ConsoleLogger.Println(string(body))

	rt := &ReceiveTag{}
	err := json.Unmarshal(body, &rt)
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
	tag := model.Tag{
		Id:         bson.NewObjectId(),
		Tagname:    rt.Tagname,
		Color:      rt.Color,
		Createtime: time.Now().Unix(),
		Updatetime: time.Now().Unix(),
		Creater:    user.Id,
		Isdelete:   false,
	}
	flag := model.InsertTag(&tag)
	if !flag {
		panic(fmt.Errorf("插入失败"))
	}
	tags, num, currentPage, pageSize, err := GetTags(bson.M{}, writer, request)
	if err != nil {
		panic(err)
	}
	responseResult := ResponseResult{
		Code:    OK,
		Message: "添加成功",
		Data: ResponseList{
			List: *tags,
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
