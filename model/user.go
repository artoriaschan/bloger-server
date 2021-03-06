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
	Registertime int64           `bson:"registertime" json:"registertime"` // 时间戳
	Description  string          `bson:"description" json:"description"`
	Articles     []bson.ObjectId `bson:"articles" json:"articles"`
	Type         int             `bson:"type" json:"type"`
	Avatar       string          `bson:"avatar" json:"avatar"`
	Freezen      bool            `bson:"freezen" json:"freezen"`
	IsDelete     bool            `bson:"isdelete" json:"isdelete"`
}

type OutPutUser struct {
	Id           bson.ObjectId `bson:"_id" json:"id"`
	Username     string        `bson:"username" json:"username"`
	Email        string        `bson:"email" json:"email"`
	Mobile       string        `bson:"mobile" json:"mobile"`
	Registertime int64         `bson:"registertime" json:"registertime"` // 时间戳
	Description  string        `bson:"description" json:"description"`
	Type         int           `bson:"type" json:"type"`
	Avatar       string        `bson:"avatar" json:"avatar"`
	Freezen      bool          `bson:"freezen" json:"freezen"`
}

func (u *User) SetPassword(password string) {
	u.Password = GeneratePasswordHash(password)
}

func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.Password
}

// 转化
func (u *User) ToOutputUser() OutPutUser {
	outputUser := new(OutPutUser)
	outputUser.Id = u.Id
	outputUser.Username = u.Username
	outputUser.Email = u.Email
	outputUser.Mobile = u.Mobile
	outputUser.Registertime = u.Registertime
	outputUser.Description = u.Description
	outputUser.Type = u.Type
	outputUser.Avatar = u.Avatar
	outputUser.Freezen = u.Freezen

	return *outputUser
}

// 根据姓名查找
func GetUserByUsername(value interface{}) (*User, bool) {
	user := new(User)
	filter := bson.M{}
	filter["username"] = value
	filter["isdelete"] = false
	flag := Find("user", filter, &user)
	return user, flag
}

// 根据邮箱
func GetUserByEmail(value interface{}) (*User, bool) {
	user := new(User)
	filter := bson.M{}
	filter["email"] = value
	filter["isdelete"] = false
	flag := Find("user", filter, &user)
	return user, flag
}

// 根据Id
func GetUserById(value interface{}) (*User, bool) {
	user := new(User)
	filter := bson.M{}
	filter["_id"] = value
	filter["isdelete"] = false
	flag := Find("user", filter, &user)
	return user, flag
}

// 获取用户列表
func GetUsers(filter bson.M, skip, limit int) (*[]OutPutUser, int, bool) {
	outPutUsers := new([]OutPutUser)
	field := bson.M{
		"_id":          1,
		"username":     1,
		"email":        1,
		"mobile":       1,
		"registertime": 1,
		"description":  1,
		"type":         1,
		"avatar":       1,
		"freezen":      1,
	}
	filter["isdelete"] = false
	countNum, flag := FindAll("user", filter, field, outPutUsers, skip, limit)
	return outPutUsers, countNum, flag
}

//删除用户
func DeleteUser(id interface{}) bool {
	selector := bson.M{"_id": bson.ObjectIdHex(id.(string))}
	data := bson.M{"$set": bson.M{"isdelete": true}}
	flag := Update("user", selector, data)
	return flag
}

// 增加用户
func InsertUser(user *User) bool {
	flag := Insert("user", user)
	return flag
}

// 冻结用户账户
func FreezeUser(id interface{}) bool {
	selector := bson.M{"_id": bson.ObjectIdHex(id.(string))}
	data := bson.M{"$set": bson.M{"freezen": true}}
	flag := Update("user", selector, data)
	return flag
}

// 解冻用户账户
func ActivteUser(id interface{}) bool {
	selector := bson.M{"_id": bson.ObjectIdHex(id.(string))}
	data := bson.M{"$set": bson.M{"freezen": false}}
	flag := Update("user", selector, data)
	return flag
}
