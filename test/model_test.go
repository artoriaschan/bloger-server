package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/artoriaschan/bloger-server/model"
	"gopkg.in/mgo.v2/bson"
)

func TestUserInsert(test *testing.T) {
	user := model.User{
		Id:           bson.NewObjectId(),
		Username:     "artorias",
		Email:        "544396118@qq.com",
		Mobile:       "18513100205",
		Registertime: time.Now().UnixNano(),
	}
	user.SetPassword("666666")

	model.InsertUser(&user)
}

func TestUserFindByUsername(test *testing.T) {
	user := model.User{}
	model.GetUserByUsername("artorias", &user)
	fmt.Println("username: ", user.Username)
	fmt.Println("password: ", user.Password)
}
func TestUserFindById(test *testing.T) {
	var objectId = bson.ObjectIdHex("5c0a6b8728c17345ccf1e1a1")
	user := new(model.User)
	user, _ = model.GetUserById(objectId)
	fmt.Println(*user)
}
