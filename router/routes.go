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
	},
	Route{
		Name:        "Login",
		Method:      []string{"Post"},
		Pattern:     "/api/login",
		HandlerFunc: controller.Login,
	},
	Route{
		Name:        "AdminLogin",
		Method:      []string{"Post"},
		Pattern:     "/api/admin/login",
		HandlerFunc: controller.AdminLogin,
	},
	Route{
		Name:        "Register",
		Method:      []string{"Post"},
		Pattern:     "/api/register",
		HandlerFunc: controller.Register,
	},
	Route{
		Name:        "CurrentAdmin",
		Method:      []string{"Get"},
		Pattern:     "/api/currentAdmin",
		HandlerFunc: controller.CurrentAdmin,
	},
}
