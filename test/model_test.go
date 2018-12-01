package test

import (
	"fmt"
	"github.com/artoriaschan/bloger-server/model"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestUserInsert(test *testing.T){
	user := model.User{
		Id: bson.NewObjectId(),
		Username: "artorias",
		Email: "544396118@qq.com",
		Mobile: "18513100205",
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