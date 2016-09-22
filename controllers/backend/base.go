package backend

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

/**
 * 前置方法用于设置默认参数、登录、权限等认证
 */
func (c *BaseController) Prepare() {
	c.Layout = "backend/layout/main.html"
	c.TplExt = "html"
	c.TplPrefix = "backend/"
}
