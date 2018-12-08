package test

import (
	"testing"

	"github.com/artoriaschan/bloger-server/utils/jwt"
)

func Test_Encode(t *testing.T) {
	jwt := jwtoken.JWT{}
	jwt.Header = jwtoken.Header{"HS256", "JWT"}
	jwt.PayLoad = jwtoken.PayLoad{"1234567890", "John Doe", "John Doe", true}
	result := jwt.Encode()
	t.Log(result)
}

func Test_Decode(t *testing.T) {
	testStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.4c9540f793ab33b13670169bdf444c1eb1c37047f18e861981e14e34587b1e04"

	jwt := jwtoken.JWT{}
	if jwt.Decode(testStr) {
		t.Log(jwt)
	} else {
		t.Error("error json content")
	}
}
