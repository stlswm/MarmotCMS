package routers

import (
	"beeframework/controllers/backend"
	"beeframework/controllers/frontend"

	"github.com/astaxie/beego"
)

func init() {
	//前台路由配置
	beego.Router("/", &frontend.DefaultController{}, "*:Index")
	beego.AutoRouter(&frontend.DefaultController{})
	beego.AutoRouter(&frontend.LoginController{})
	//后台路由配置
	bns := beego.NewNamespace("/backend",
		beego.NSAutoRouter(&backend.DefaultController{}),
		beego.NSAutoRouter(&backend.LoginController{}),
		beego.NSAutoRouter(&backend.CategoryController{}),
	)
	beego.AddNamespace(bns)
}
