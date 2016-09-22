package frontend

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (c *DefaultController) Index() {
	c.Ctx.WriteString("frontend/default/index")
}
