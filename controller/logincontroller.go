package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/artoriaschan/bloger-server/utils/logging"
	"github.com/go-session/session"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/artoriaschan/bloger-server/model"
	"gopkg.in/mgo.v2/bson"
)

func userLoginHandle(username, password string) ([]byte, *model.User, bool) {
	var result []byte
	var isSuccess bool
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
		isSuccess = false
		return result, nil, isSuccess
	}
	flag := user.CheckPassword(password)
	if flag {
		responseResult := ResponseResult{
			Code:    OK,
			Message: "查询成功",
			Data:    user,
		}
		result = responseResult.ToJson()
		isSuccess = true
	} else {
		responseResult := ResponseResult{
			Code:    WrongParams,
			Message: "邮箱/密码输入错误",
			Data:    nil,
		}
		result = responseResult.ToJson()
		isSuccess = false
	}
	return result, &user, isSuccess
}
//登录
func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	//writer.WriteHeader(http.StatusOK)
	fmt.Println(request.Cookies())
	var email string
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
		password = request.Form.Get("password")
	}
	result, user, isSuccess := userLoginHandle(email, password)
	if isSuccess {
		store, err := session.Start(context.Background(), writer, request)
		if err != nil {
			panic(err)
		}
		store.Set("loginUser", user)
		// 写入cookie
		COOKIE_MAX_MAX_AGE := time.Hour * 24 / time.Second   // 单位：秒。
		maxAge := int(COOKIE_MAX_MAX_AGE)
		cookie := http.Cookie{
			Name:   "uid",
			Value:  user.Id.Hex(),
			Path: "/",
			HttpOnly:true,
			MaxAge: maxAge,
		}
		http.SetCookie(writer, &cookie)
	}
	writer.Write(result)

	// 日志处理
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
}
// 后台登录
func AdminLogin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println(request.Cookies())
	var email string
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
		password = request.Form.Get("password")
	}
	result, user, isSuccess := userLoginHandle(email, password)
	if isSuccess {
		store, err := session.Start(context.Background(), writer, request)
		if err != nil {
			panic(err)
		}
		store.Set("loginUser", user)
		// 写入cookie
		COOKIE_MAX_MAX_AGE := time.Hour * 24 / time.Second   // 单位：秒。
		maxAge := int(COOKIE_MAX_MAX_AGE)
		cookie := http.Cookie{
			Name:   "uid",
			Value:  user.Id.Hex(),
			Path: "/",
			HttpOnly:true,
			MaxAge: maxAge,
		}
		http.SetCookie(writer, &cookie)
	}
	writer.Write(result)

	// 日志处理
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
				Code:    WrongParams,
				Message: "邮箱格式不正确",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}

		flag := model.GetUserByEmail(email, &model.User{})
		if flag {
			responseResult = ResponseResult{
				Code:    WrongParams,
				Message: "该邮箱已注册",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	} else {
		responseResult = ResponseResult{
			Code:    WrongParams,
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
				Code:    WrongParams,
				Message: "用户名填写错误",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	} else {
		if len(username) > 12 || len(username) < 5 {
			responseResult = ResponseResult{
				Code:    WrongParams,
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
				Code:    WrongParams,
				Message: "密码填写错误",
				Data:    nil,
			}
			result = responseResult.ToJson()
			return result
		}
	} else {
		responseResult = ResponseResult{
			Code:    WrongParams,
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
				Code:    WrongParams,
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
// 注册
func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
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
