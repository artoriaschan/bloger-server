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
		Name:        "Register",
		Method:      []string{"Post"},
		Pattern:     "/api/register",
		HandlerFunc: controller.Register,
	},
	Route{
		Name:        "Login",
		Method:      []string{"Post"},
		Pattern:     "/api/login",
		HandlerFunc: controller.Login,
	},
	Route{
		Name:        "CurrentUser",
		Method:      []string{"Get"},
		Pattern:     "/api/current/user",
		HandlerFunc: controller.CurrentUser,
	},
	Route{
		Name:        "AdminLogin",
		Method:      []string{"Post"},
		Pattern:     "/api/admin/login",
		HandlerFunc: controller.AdminLogin,
	},
	Route{
		Name:        "CurrentAdmin",
		Method:      []string{"Get"},
		Pattern:     "/api/currentAdmin",
		HandlerFunc: controller.CurrentAdmin,
	},
	Route{
		Name:        "GetUsers",
		Method:      []string{"Get"},
		Pattern:     "/api/admin/users",
		HandlerFunc: controller.GetUsers,
	},
	Route{
		Name:        "DeleteUser",
		Method:      []string{"Get"},
		Pattern:     "/api/admin/users/delete/{userId}",
		HandlerFunc: controller.DeleteUser,
	},
	Route{
		Name:        "FreezeUser",
		Method:      []string{"Get"},
		Pattern:     "/api/admin/users/freeze/{userId}",
		HandlerFunc: controller.FreezeUser,
	},
	Route{
		Name:        "ActivteUser",
		Method:      []string{"Get"},
		Pattern:     "/api/admin/users/activite/{userId}",
		HandlerFunc: controller.ActivteUser,
	},
	// 文章相关
	Route{
		Name:        "AddArticle",
		Method:      []string{"Post"},
		Pattern:     "/api/article/post",
		HandlerFunc: controller.AddArticle,
	},
	// 分类相关
	Route{
		Name:        "AddCategory",
		Method:      []string{"Post"},
		Pattern:     "/api/cate/post",
		HandlerFunc: controller.AddCategory,
	},
	Route{
		Name:        "QueryCategories",
		Method:      []string{"Get"},
		Pattern:     "/api/cates",
		HandlerFunc: controller.QueryCategories,
	},
	Route{
		Name:        "ModifyCategory",
		Method:      []string{"Post"},
		Pattern:     "/api/cates/update",
		HandlerFunc: controller.ModifyCategory,
	},
	Route{
		Name:        "DeleteCategory",
		Method:      []string{"Get"},
		Pattern:     "/api/cate/delete/{cateId}",
		HandlerFunc: controller.DeleteCategory,
	},
	// 标签相关
	Route{
		Name:        "AddTag",
		Method:      []string{"Post"},
		Pattern:     "/api/tag/post",
		HandlerFunc: controller.AddTag,
	},
	Route{
		Name:        "QueryTags",
		Method:      []string{"Get"},
		Pattern:     "/api/tags",
		HandlerFunc: controller.QueryTags,
	},
}
