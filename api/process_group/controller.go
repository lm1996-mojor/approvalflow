package process_group

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

// GetProcessGroupsPage
//
// @Summary 根据条件获取流程分组列表
// @Description 根据条件获取流程分组列表
// @Tags　流程分组管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/groups/list [get]
func (c *Controller) GetProcessGroupsPage() rest.Result {
	var queryMdl QueryPageMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(queryMdl.AppCode) <= 0 {
		return rest.FailCustom(400, "请确定应用", rest.ERROR)
	}
	return c.Svc.ObtainProcessGroupsList(queryMdl)
}

// GetProcessGroupsBy
//
// @Summary 获取单个流程分组的信息
// @Description 获取单个流程分组信息
// @Tags　流程分组管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "流程分组id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/groups/{id} [get]
func (c *Controller) GetProcessGroupsBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择流程分组", rest.ERROR)
	}
	return c.Svc.ObtainProcessGroupsInfo(id)
}

// PostProcessGroups
//
// @Summary 新增流程分组
// @Description 新增流程分组
// @Tags　流程分组管理
// @Accept application/json
// @Produce application/json
// @Param object model.ProcessGroups
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/groups [get]
func (c *Controller) PostProcessGroups() rest.Result {
	var processGroupInfo model.ProcessGroups
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddProcessGroups(processGroupInfo)
}

// PutProcessGroups
//
// @Summary 根据id修改流程分组信息
// @Description 根据id修改流程分组信息
// @Tags　流程分组管理
// @Accept application/json
// @Produce application/json
// @Param object model.ProcessGroups
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/groups [get]
func (c *Controller) PutProcessGroups() rest.Result {
	var processGroupInfo model.ProcessGroups
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeProcessGroups(processGroupInfo)
}

// DeleteProcessGroups
//
// @Summary 删除流程分组
// @Description 删除流程分组
// @Tags　流程分组管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/process/groups [get]
func (c *Controller) DeleteProcessGroups() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择流程分组", rest.ERROR)
	}
	return c.Svc.DropProcessGroups(delMdl.Ids)
}

// GetProcessGroupsAllList
//
// @Summary 获取全部流程分组列表
//
// @Description 获取全部流程分组列表
//
// @Group　流程分组管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/process/groups/all/list
func (c *Controller) GetProcessGroupsAllList() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(queryMdl.AppCode) <= 0 {
		return rest.FailCustom(400, "请确定应用", rest.ERROR)
	}
	return c.Svc.ObtainProcessGroupsAllList(queryMdl)
}
