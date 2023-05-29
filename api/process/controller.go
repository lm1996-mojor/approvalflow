package process

import (
	"five.com/lk_flow/api/process_group"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest"
	"five.com/technical_center/core_library.git/rest/req"
	"github.com/kataras/iris/v12"
)

type Controller struct {
	Ctx iris.Context
	Svc Service
}

func NewController(repository Repository, groupRepo process_group.Repository) *Controller {
	return &Controller{Svc: NewService(repository, groupRepo)}
}

// GetProcessPage
//
// @Summary 根据条件获取流程列表
// @Description 根据条件获取流程列表
// @Tags　流程管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/list [get]
func (c *Controller) GetProcessPage() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainProcessPage(queryMdl)
}

// GetProcessBy
//
// @Summary 获取单个流程的信息
// @Description 获取单个流程信息
// @Tags　流程管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "流程id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/{id} [get]
func (c *Controller) GetProcessBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择流程", rest.ERROR)
	}
	return c.Svc.ObtainProcessInfo(id)
}

// PostProcess
//
// @Summary 新增流程主体的信息
// @Description 新增流程主体的信息
// @Tags　流程管理
// @Accept application/json
// @Produce application/json
// @Param object model.Process
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process [get]
func (c *Controller) PostProcess() rest.Result {
	// TODO: 创建完成后需要返回ID
	var processInfo model.Process
	if err := c.Ctx.ReadJSON(&processInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddProcess(processInfo)
}

// PutProcess
//
// @Summary 根据id修改流程信息
// @Description 根据id修改流程信息
// @Tags　流程管理
// @Accept application/json
// @Produce application/json
// @Param object model.Process
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process [get]
func (c *Controller) PutProcess() rest.Result {
	var processInfo model.Process
	if err := c.Ctx.ReadJSON(&processInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeProcess(processInfo)
}

// DeleteProcess
//
// @Summary 删除流程
// @Description 删除流程
// @Tags　流程管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process [get]
func (c *Controller) DeleteProcess() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择流程", rest.ERROR)
	}
	return c.Svc.DropProcess(delMdl.Ids)
}

// PutReleaseProcess
//
// @Summary 发布指定流程
// @Description 发布指定流程
// @Tags　流程管理
// @Accept application/json
// @Produce application/json
// @Param object model.Process
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/release/process [get]
func (c *Controller) PutReleaseProcess(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择流程", rest.ERROR)
	}
	return c.Svc.ReleaseProcess(id)
}
