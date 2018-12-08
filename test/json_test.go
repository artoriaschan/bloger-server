package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/artoriaschan/bloger-server/controller"
)

func Test_json(t *testing.T) {
	var body = []byte("{\"email\":\"chenzheng04@58ganji.com\",\"password\":\"666666\"}")
	bodyStruct := controller.AdminLoginBody{}
	err := json.Unmarshal(body, &bodyStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bodyStruct.Email)
}
