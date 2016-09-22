package backend

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	c.TplExt = "html"
	c.TplPrefix = "backend/"
}

/**
 * 后台管理员登录
 */
func (c *LoginController) Login() {
	if c.Ctx.Request.Method == "GET" {
		c.Render()
	} else {
		c.Redirect("/backend/default/index", 301)
	}
}

/**
 * 后台管理员登录登出
 */
func (c *LoginController) Logout() {
	c.Redirect("/backend/login/login", 301)
}
