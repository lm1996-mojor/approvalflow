package panel_point

import (
	"five.com/lk_flow/api/process"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/rest/req"
	"github.com/kataras/iris/v12"
)

type Controller struct {
	Ctx iris.Context
	Svc Service
}

func NewController(repository Repository, processRepo process.Repository) *Controller {
	return &Controller{Svc: NewService(repository, processRepo)}
}

// GetPanelPointList
//
// @Summary 根据条件获取流程节点列表
// @Description 根据条件获取流程节点列表
// @Tags　流程节点管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/panel/point/list [get]
func (c *Controller) GetPanelPointList() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainPanelPointList(queryMdl)
}

// GetPanelPointBy
//
// @Summary 获取单个流程节点的信息
// @Description 获取单个流程节点信息
// @Tags　流程节点管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "流程节点id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/panel/point/{id} [get]
func (c *Controller) GetPanelPointBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择流程节点", rest.ERROR)
	}
	return c.Svc.ObtainPanelPointInfo(id)
}

// PostPanelPoint
//
// @Summary 新增流程接节点信息
// @Description 新增流程接节点信息
// @Tags　流程节点管理
// @Accept application/json
// @Produce application/json
// @Param object model.PanelPoint
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/panel/point [get]
func (c *Controller) PostPanelPoint() rest.Result {
	var panelPointDetail PanelPointDetail
	if err := c.Ctx.ReadJSON(&panelPointDetail); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddPanelPoint(panelPointDetail)
}

// PutPanelPoint
//
// @Summary 根据id修改流程节点信息
// @Description 根据id修改流程节点信息
// @Tags　流程节点管理
// @Accept application/json
// @Produce application/json
// @Param object model.PanelPoint
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/panel/point [get]
func (c *Controller) PutPanelPoint() rest.Result {
	var processInfo model.PanelPoint
	if err := c.Ctx.ReadJSON(&processInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangePanelPoint(processInfo)
}

// DeletePanelPoint
//
// @Summary 删除流程节点
// @Description 删除流程节点
// @Tags　流程节点管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/panel/point [get]
func (c *Controller) DeletePanelPoint() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择流程节点", rest.ERROR)
	}
	return c.Svc.DropPanelPoint(delMdl.Ids)
}
