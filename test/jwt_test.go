package test

import (
	"github.com/artoriaschan/bloger-server/utils"
	"testing"
)

func Test_Encode(t *testing.T) {
	jwt := utils.JWT{}
	jwt.Header = utils.Header{"HS256","JWT"}
	jwt.PayLoad = utils.PayLoad{"1234567890","John Doe",true}
	result := jwt.Encode()
	t.Log(result)
}

func Test_Decode(t *testing.T) {
	testStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.4c9540f793ab33b13670169bdf444c1eb1c37047f18e861981e14e34587b1e04"

	jwt := utils.JWT{}
	if jwt.Decode(testStr) {
		t.Log(jwt)
	} else {
		t.Error("error json content")
	}
}
