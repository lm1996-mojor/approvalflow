package comparison_value

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

// GetComparisonValueList
//
// @Summary 根据条件获取比较值信息列表
//
// @Description 根据条件获取比较值信息列表
//
// @Tags　比较值信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object QueryListMdl
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/value/list [get]
func (c *Controller) GetComparisonValueList() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if queryMdl.CondDetailInfoId <= 0 {
		return rest.FailCustom(400, "请选择一个条件详细信息", rest.ERROR)
	}
	return c.Svc.ObtainComparisonValueList(queryMdl)
}

// GetComparisonValueBy
//
// @Summary 获取单个比较值信息的信息
//
// @Description 获取单个比较值信息信息
//
// @Tags　比较值信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param id path int true "比较值信息id"
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/value/{id} [get]
func (c *Controller) GetComparisonValueBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择比较值信息", rest.ERROR)
	}
	return c.Svc.ObtainComparisonValueInfo(id)
}

// PostComparisonValue
//
// @Summary 新增比较值信息
//
// @Description 新增比较值信息
//
// @Tags　比较值信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object model.ComparisonValue
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/value [post]
func (c *Controller) PostComparisonValue() rest.Result {
	var processGroupInfo model.ComparisonValue
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddComparisonValue(processGroupInfo)
}

// PutComparisonValue
//
// @Summary 根据id修改比较值信息信息
//
// @Description 根据id修改比较值信息信息
//
// @Tags　比较值信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object model.ComparisonValue
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/value [put]
func (c *Controller) PutComparisonValue() rest.Result {
	var processGroupInfo model.ComparisonValue
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeComparisonValue(processGroupInfo)
}

// DeleteComparisonValue
//
// @Summary 删除比较值信息
//
// @Description 删除比较值信息
//
// @Tags　比较值信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object req.DelModel
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/value [delete]
func (c *Controller) DeleteComparisonValue() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择比较值信息", rest.ERROR)
	}
	return c.Svc.DropComparisonValue(delMdl.Ids)
}
