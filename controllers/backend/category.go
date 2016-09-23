package backend

//"MarmotCMS/models"

type CategoryController struct {
	BaseController
}

/**/
func (c *CategoryController) Index() {
	c.Data["nav_focus"] = 1
	c.Render()
}

func (c *CategoryController) AddNormal() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["nav_focus"] = 2
		c.Render()
	} else {

	}
}
func (c *CategoryController) AddPage() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["nav_focus"] = 3
		c.Render()
	} else {

	}
}
func (c *CategoryController) AddOuter() {

	if c.Ctx.Request.Method == "GET" {
		c.Data["nav_focus"] = 4
		c.Render()
	} else {

	}
}
