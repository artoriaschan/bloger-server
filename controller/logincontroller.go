package controller

import (
	"encoding/json"
	"fmt"
	"github.com/artoriaschan/bloger-server/utils/logging"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/artoriaschan/bloger-server/model"
	"gopkg.in/mgo.v2/bson"
)

func userLoginHandle(username, password string) ([]byte, *model.User) {
	var result []byte
	user := model.User{}
	//hasUser := model.GetUserByUsername(username, &user)
	hasUser := model.GetUserByEmail(username, &user)
	if !hasUser {
		responseResult := ResponseResult{
			Code:    NoRegister,
			Message: "该邮箱没有注册",
			Data:    nil,
		}
		result = responseResult.ToJson()
		return result, nil
	}
	flag := user.CheckPassword(password)
	if flag {
		responseResult := ResponseResult{
			Code:    OK,
			Message: "查询成功",
			Data:    user,
		}
		result = responseResult.ToJson()
	} else {
		responseResult := ResponseResult{
			Code:    WrongPassword,
			Message: "邮箱/密码输入错误",
			Data:    nil,
		}
		result = responseResult.ToJson()
	}
	return result, &user
}
func Login(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	var email string
	var password string
	if request.Method == http.MethodGet {
		query := request.URL.RawQuery
		queryMap, _ := url.ParseQuery(query)
		fmt.Println(queryMap)
		email = queryMap["email"][0]
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
		email = request.Form.Get("email")
		password = request.Form.Get("password")
	}
	result, user := userLoginHandle(email, password)
	// 写入cookie
	cookieValue := "username=" + user.Username + "&email=" + user.Email
	// COOKIE_MAX_MAX_AGE := 30 * time.Hour * 24 / time.Second // 单位：秒。
	// maxAge := int(COOKIE_MAX_MAX_AGE)
	cookie := http.Cookie{
		Name:   "onebitcode",
		Value:  cookieValue,
		Path:   "/",
		MaxAge: -1,
		Domain: "localhost", //域名
	}
	http.SetCookie(writer, &cookie)
	header, _ := json.Marshal(request.Header)
	userJson, _ := json.Marshal(user)
	access := logging.AccessLoggerFormat{
		IP: request.RemoteAddr,
		Header: string(header),
		UserAgent: request.UserAgent(),
		Extend: string(userJson),
	}
	accessJson,_ := json.Marshal(access)
	ConsoleLogger.Println(string(accessJson))
	AccessLogger.Println(string(accessJson))
	writer.Write(result)
}

func userRegisterHandle(email, username, mobile, password string, user *model.User) []byte {
	var responseResult ResponseResult
	var result []byte

	// email validator
	if email != "" {
		emailExp := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
		emailExpResult := emailExp.FindAllStringSubmatch(email, -1)
		if emailExpResult == nil {
			responseResult = ResponseResult{
				Code:    WrongFormatEmail,
				Message: "邮箱格式不正确",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}

		flag := model.GetUserByEmail(email, &model.User{})
		if flag {
			responseResult = ResponseResult{
				Code:    ExsitedEmail,
				Message: "该邮箱已注册",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	} else {
		responseResult = ResponseResult{
			Code:    EmptyEmail,
			Message: "邮箱不能为空",
			Data:    nil,
		}
		result = responseResult.ToJson()
		return result
	}

	// username validator
	if username != "" {
		if len(username) > 12 || len(username) < 5 {
			responseResult = ResponseResult{
				Code:    WrongFormatUsername,
				Message: "用户名填写错误",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	} else {
		if len(username) > 12 || len(username) < 5 {
			responseResult = ResponseResult{
				Code:    EmptyUsername,
				Message: "用户名不能为空",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	}
	// password validator
	if password != "" {
		if len(password) < 6 || len(password) > 18 {
			responseResult = ResponseResult{
				Code:    WrongFormatPassword,
				Message: "密码填写错误",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	} else {
		responseResult = ResponseResult{
			Code:    EmpeyPassword,
			Message: "密码不能为空",
			Data:    nil,
		}
		result = responseResult.ToJson()
		return result
	}
	// mobile validator
	if mobile != "" {
		mobileExp := regexp.MustCompile(`^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`)
		mobileExpResult := mobileExp.FindAllStringSubmatch(email, -1)
		if mobileExpResult == nil {
			responseResult = ResponseResult{
				Code:    WrongFormatMobile,
				Message: "手机格式不正确",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	}
	user = &model.User{
		Id:           bson.NewObjectId(),
		Username:     username,
		Email:        email,
		Mobile:       mobile,
		Registertime: time.Now().UnixNano(),
	}
	user.SetPassword(password)

	flag := model.InsertUser(user)

	if flag {
		responseResult = ResponseResult{
			Code:    OK,
			Message: "注册成功",
			Data:    user,
		}
		result = responseResult.ToJson()
	} else {
		responseResult = ResponseResult{
			Code:    NoRegister,
			Message: "注册失败",
			Data:    nil,
		}
		result = responseResult.ToJson()
	}
	return result
}

func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	var result []byte
	var email string
	var username string
	var mobile string
	var password string

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

		email = request.Form.Get("email")
		username = request.Form.Get("username")
		mobile = request.Form.Get("mobile")
		password = request.Form.Get("password")

		user := model.User{}
		result = userRegisterHandle(email, username, mobile, password, &user)
	}

	writer.Write(result)
}
