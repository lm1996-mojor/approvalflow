package condition_detail_groups

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

// GetConditionDetailGroupsPage
//
// @Summary 根据条件获取条件详细信息分组列表
// @Description 根据条件获取条件详细信息分组列表
// @Tags　条件详细信息分组管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups/list [get]
func (c *Controller) GetConditionDetailGroupsPage() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if queryMdl.ConditionInfoId <= 0 {
		return rest.FailCustom(500, "请确认条件主体", rest.ERROR)
	}
	return c.Svc.ObtainConditionDetailGroupsList(queryMdl)
}

// GetConditionDetailGroupsBy
//
// @Summary 获取单个条件详细信息分组的信息
// @Description 获取单个条件详细信息分组信息
// @Tags　条件详细信息分组管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "条件详细信息分组id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups/{id} [get]
func (c *Controller) GetConditionDetailGroupsBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择条件详细信息分组", rest.ERROR)
	}
	return c.Svc.ObtainConditionDetailGroupsInfo(id)
}

// PostConditionDetailGroups
//
// @Summary 新增条件详细信息分组
// @Description 新增条件详细信息分组
// @Tags　条件详细信息分组管理
// @Accept application/json
// @Produce application/json
// @Param object model.ConditionDetailGroups
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups [post]
func (c *Controller) PostConditionDetailGroups() rest.Result {
	var processGroupInfo model.ConditionDetailGroups
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddConditionDetailGroups(processGroupInfo)
}

// PutConditionDetailGroups
//
// @Summary 根据id修改条件详细信息分组信息
// @Description 根据id修改条件详细信息分组信息
// @Tags　条件详细信息分组管理
// @Accept application/json
// @Produce application/json
// @Param object model.ConditionDetailGroups
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups [put]
func (c *Controller) PutConditionDetailGroups() rest.Result {
	var processGroupInfo model.ConditionDetailGroups
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeConditionDetailGroups(processGroupInfo)
}

// DeleteConditionDetailGroups
//
// @Summary 删除条件详细信息分组
// @Description 删除条件详细信息分组
// @Tags　条件详细信息分组管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups [delete]
func (c *Controller) DeleteConditionDetailGroups() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择条件详细信息分组", rest.ERROR)
	}
	return c.Svc.DropConditionDetailGroups(delMdl.Ids)
}
