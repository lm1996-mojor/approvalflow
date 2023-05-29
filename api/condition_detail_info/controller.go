package condition_detail_info

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

// GetConditionDetailInfoPage
//
// @Summary 根据条件获取条件详细信息列表
// @Description 根据条件获取条件详细信息列表
// @Tags　条件详细信息管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups/list [get]
func (c *Controller) GetConditionDetailInfoPage() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if queryMdl.GroupsId <= 0 {
		return rest.FailCustom(400, "请确认条件详情分组", rest.ERROR)
	}
	return c.Svc.ObtainConditionDetailInfoList(queryMdl)
}

// GetConditionDetailInfoBy
//
// @Summary 获取单个条件详细信息的信息
// @Description 获取单个条件详细信息信息
// @Tags　条件详细信息管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "条件详细信息id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups/{id} [get]
func (c *Controller) GetConditionDetailInfoBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择条件详细信息", rest.ERROR)
	}
	return c.Svc.ObtainConditionDetailInfoInfo(id)
}

// PostConditionDetailInfo
//
// @Summary 新增条件详细信息
// @Description 新增条件详细信息
// @Tags　条件详细信息管理
// @Accept application/json
// @Produce application/json
// @Param object model.ConditionDetailInfo
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups [post]
func (c *Controller) PostConditionDetailInfo() rest.Result {
	var conditionDetailInfo model.ConditionDetailInfo
	if err := c.Ctx.ReadJSON(&conditionDetailInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if conditionDetailInfo.GroupsId <= 0 {
		return rest.FailCustom(400, "请确认分组", rest.ERROR)
	}
	return c.Svc.AddConditionDetailInfo(conditionDetailInfo)
}

// PutConditionDetailInfo
//
// @Summary 根据id修改条件详细信息信息
// @Description 根据id修改条件详细信息信息
// @Tags　条件详细信息管理
// @Accept application/json
// @Produce application/json
// @Param object model.ConditionDetailInfo
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups [put]
func (c *Controller) PutConditionDetailInfo() rest.Result {
	var processGroupInfo model.ConditionDetailInfo
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeConditionDetailInfo(processGroupInfo)
}

// DeleteConditionDetailInfo
//
// @Summary 删除条件详细信息
// @Description 删除条件详细信息
// @Tags　条件详细信息管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/condition/detail/groups [delete]
func (c *Controller) DeleteConditionDetailInfo() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择条件详细信息", rest.ERROR)
	}
	return c.Svc.DropConditionDetailInfo(delMdl.Ids)
}
