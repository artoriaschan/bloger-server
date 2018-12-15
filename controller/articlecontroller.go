package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/artoriaschan/bloger-server/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type ReceiveArticle struct {
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	Num             int      `json:"num"`
	Desc            string   `json:"desc"`
	Cover           string   `json:"cover"`
	Keywords        string   `json:"keywords"`
	ArticleCates    []string `json:"articleCates"`
	ArticleTags     []string `json:"articleTags"`
	ArticleOperator int      `json:"articleOperator"`
	ArticleType     int      `json:"articleType"`
}

// /api/article/post
func AddArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(request.Body)

	re := &ReceiveArticle{}
	err := json.Unmarshal(body, &re)
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
	// TODO 判断是否为空
	if sess.Get("loginAdmin") == nil {
		panic(fmt.Errorf("未登录"))
	}
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
	tags := []bson.ObjectId{}
	cates := []bson.ObjectId{}
	for _, v := range re.ArticleTags {
		tags = append(tags, bson.ObjectIdHex(v))
	}
	for _, v := range re.ArticleCates {
		cates = append(tags, bson.ObjectIdHex(v))
	}
	article := model.Article{
		Id:          bson.NewObjectId(),
		Createtime:  time.Now().Unix(),
		Title:       re.Title,
		Tags:        tags,
		Categories:  cates,
		Content:     re.Content,
		Description: re.Desc,
		Author:      admin.Id,
		Updatetime:  time.Now().Unix(),
		Meta: model.Meta{
			Watch: 0,
			Like:  0,
			Num:   re.Num,
		},
		Cover:    re.Cover,
		Keywords: re.Keywords,
		Type:     re.ArticleType,
		Operator: re.ArticleOperator,
		Isdelete: false,
	}
	// 将文章ID插入到User中
	ok := model.InsertArticle(&article)
	if !ok {
		fmt.Errorf("插入失败")
	}
	responseResult := ResponseResult{
		Code:    OK,
		Message: "添加成功",
		Data:    nil,
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}

// /api/article/{articleId}
func QueryArticleDetail(writer http.ResponseWriter, request *http.Request) {
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
	articleId := vars["articleId"]
	article, ok := model.FindArticleById(articleId)
	if !ok {
		fmt.Errorf("没有找到文章")
	}
	responseResult := ResponseResult{
		Code:    OK,
		Message: "查询成功",
		Data:    *article,
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}
