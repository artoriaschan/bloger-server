package model

import (
	"gopkg.in/mgo.v2/bson"
)

/*
{
    id ObjectId,    *
    useraname string,   *
    passowrd string,    *
    mobile string,
    registertime timestemp, *
    description string,
    articles []ObjectId,
    type int,   *   // 1: user, 2: admin
    avatar string,
    freezen bool,   *
    isdelete bool   *
}
*/
type User struct {
	Id           bson.ObjectId   `bson:"_id" json:"id"`
	Username     string          `bson:"username" json:"username"`
	Email        string          `bson:"email" json:"email"`
	Password     string          `bson:"password" json:"password"`
	Mobile       string          `bson:"mobile" json:"mobile"`
	Registertime int64           `bson:"registertime" json:"registerTime"` // 时间戳
	Description  string          `bson:"description" json:"description"`
	Articles     []bson.ObjectId `bson:"articles" json:"articles"`
	Type         int             `bson:"type" json:"type"`
	Avatar       string          `bson:"avatar" json:"avatar"`
	Freezen      bool            `bson:"freezen" json:"freezen"`
	IsDelete     bool            `bson:"isdelete" json:"isdelete"`
}

func (u *User) SetPassword(password string) {
	u.Password = GeneratePasswordHash(password)
}

func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.Password
}

// 根据姓名查找
func GetUserByUsername(value interface{}) (*User, bool) {
	user := new(User)
	flag := Find("user", bson.M{"username": value}, &user)
	return user, flag
}

// 根据邮箱
func GetUserByEmail(value interface{}) (*User, bool) {
	user := new(User)
	flag := Find("user", bson.M{"email": value}, &user)
	return user, flag
}

// 根据Id
func GetUserById(value interface{}) (*User, bool) {
	user := new(User)
	flag := Find("user", bson.M{"_id": value}, &user)
	return user, flag
}

// 增加用户
func InsertUser(user *User) bool {
	flag := Insert("user", user)
	return flag
}
