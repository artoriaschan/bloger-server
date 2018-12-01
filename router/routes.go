package router

import (
	"net/http"

	"github.com/artoriaschan/bloger-server/controller"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      []string{"Get"},
		Pattern:     "/",
		HandlerFunc: controller.IndexTemplate,
		// func(writer http.ResponseWriter, request *http.Request) {
		// 	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 	writer.WriteHeader(http.StatusOK)
		// 	writer.Write([]byte("{'path':'/','result':'ok'}"))
		// 	//// Case 1: w.Write byte
		// 	//w.Write([]byte("Hello World"))
		// 	//// Case 2: fmt.Fprintf
		// 	//fmt.Fprintf(w, "Hello World")
		// 	//// Case 3: io.Write
		// 	//io.WriteString(w, "Hello World")
		// },
	},
	Route{
		Name:        "Login",
		Method:      []string{"Get", "Post"},
		Pattern:     "/api/login",
		HandlerFunc: controller.Login,
	},
	Route{
		Name:        "Register",
		Method:      []string{"Post"},
		Pattern:     "/api/register",
		HandlerFunc: controller.Register,
	},
}
