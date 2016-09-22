package frontend

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

/*
 *用户登录
 */
func (c *LoginController) Login() {
	if c.Ctx.Request.Method == "GET" {
		c.Ctx.WriteString("frontend login")
	} else {

	}
}

/*
 *用户登出
 */
func (c *LoginController) Logout() {
	c.Redirect("/", 301)
}
