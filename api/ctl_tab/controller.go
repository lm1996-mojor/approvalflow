package approval_rule_go

import (
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/rest/req"
	"github.com/kataras/iris/v12"
)

type Controller struct {
	Ctx iris.Context
	Svc Service
}

func NewController(repository Repository) *Controller {
	return &Controller{Svc: NewService(repository)}
}

// GetCtlTabPage
//
// @Summary 根据条件获取条件要素（控件标签）列表
// @Description 根据条件获取条件要素（控件标签）列表
// @Tags　条件要素（控件标签）管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/ctl/tab/page [get]
func (c *Controller) GetCtlTabPage() rest.Result {
	var queryMdl QueryPageMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainCtlTabPage(queryMdl)
}

// GetCtlTabBy
//
// @Summary 获取单个条件要素（控件标签）的信息
// @Description 获取单个条件要素（控件标签）信息
// @Tags　条件要素（控件标签）管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "条件要素（控件标签）id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/ctl/tab/{id} [get]
func (c *Controller) GetCtlTabBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择条件要素（控件标签）", rest.ERROR)
	}
	return c.Svc.ObtainCtlTabInfo(id)
}

// PostCtlTab
//
// @Summary 新增条件要素（控件标签）
// @Description 新增条件要素（控件标签）
// @Tags　条件要素（控件标签）管理
// @Accept application/json
// @Produce application/json
// @Param object model.CtlTab
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/ctl/tab [get]
func (c *Controller) PostCtlTab() rest.Result {
	var info model.CtlTab
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddCtlTab(info)
}

// PutCtlTab
//
// @Summary 根据id修改条件要素（控件标签）信息
// @Description 根据id修改条件要素（控件标签）信息
// @Tags　条件要素（控件标签）管理
// @Accept application/json
// @Produce application/json
// @Param object model.CtlTab
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/ctl/tab [get]
func (c *Controller) PutCtlTab() rest.Result {
	var info model.CtlTab
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeCtlTab(info)
}

// DeleteCtlTab
//
// @Summary 删除条件要素（控件标签）
// @Description 删除条件要素（控件标签）
// @Tags　条件要素（控件标签）管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/ctl/tab [get]
func (c *Controller) DeleteCtlTab() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择条件要素（控件标签）", rest.ERROR)
	}
	return c.Svc.DropCtlTab(delMdl.Ids)
}

// GetCtlTabAllList
//
// @Summary 获取全部参与者形式
// @Description 获取全部参与者形式
// @Tags　参与者形式管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/panel/point/page [get]
func (c *Controller) GetCtlTabAllList() rest.Result {
	return c.Svc.ObtainCtlTabAllList()
}
