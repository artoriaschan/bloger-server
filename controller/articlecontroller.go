package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ReceiveArticle struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	Num             int    `json:"num"`
	Author          string `json:"author"`
	Desc            string `json:"desc"`
	Cover           string `json:"cover"`
	Keywords        string `json:"keywords"`
	ArticleCates    []int  `json:"articleCates"`
	ArticleTags     []int  `json:"articleTags"`
	ArticleOperator int    `json:"articleOperator"`
	ArticleType     int    `json:"articleType"`
}

// /api/article/post
func AddArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, _ := ioutil.ReadAll(request.Body)
	ConsoleLogger.Println(string(body))
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
	responseResult := ResponseResult{
		Code:    OK,
		Message: "添加成功",
		Data:    *re,
	}
	result := responseResult.ToJson()
	writer.Write(result)
	return
}
