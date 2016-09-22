package backend

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Index() {
	c.Render()
}
