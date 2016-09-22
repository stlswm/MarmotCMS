package backend

type DefaultController struct {
	BaseController
}

/**
 * 栏目首页
 */
func (c *DefaultController) Index() {
	c.Render()
}

/**
 * 添加普通栏目
 */
func (c *DefaultController) AddNoraml() {}

/**
 * 添加单页网
 */
func (c *DefaultController) AddPage() {}

/**
 * 添加外部栏目
 */
func (c *DefaultController) AddOut() {

}

/**
 * 修改普通栏目
 */
func (c *DefaultController) EditNoraml() {}

/**
 * 修改单页网
 */
func (c *DefaultController) EditPage() {}

/**
 * 修改外部栏目
 */
func (c *DefaultController) EditOut() {

}

/**
 * 删除栏目
 */
func (c *DefaultController) Delete() {

}

/**
 * 栏目排序
 */
func (c *DefaultController) Order() {

}

/**
 * 栏目批量移动
 */
func (c *DefaultController) BatchMove() {

}
