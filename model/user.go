package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id",json:"id"`
	Username string `bson:"username",json:"username"`
	Email string `bson:"email",json:"email"`
	Password string `bson:"password",json:"password"`
	Mobile string `bson:"mobile",json:"mobile"`
	Registertime int64 `bson:"register_time",json:"registerTime"`	// 时间戳
}

func (u *User) SetPassword (password string) {
	u.Password = GeneratePasswordHash(password)
}

func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.Password
}

// 根据姓名查找
func GetUserByUsername(value string, user *User) bool{
	flag := Find("user",bson.M{"username" : value}, user)
	return flag
}

func GetUserByEmail(value string, user *User) bool {
	flag := Find("user",bson.M{"email" : value}, user)
	return flag
}

// 增加用户
func InsertUser(user *User) bool{
	flag:= Insert("user", user)
	return flag
}