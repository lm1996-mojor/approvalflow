package comparison_operators

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

// GetComparisonOperatorsList
//
// @Summary 根据条件获取比较符信息分页列表
//
// @Description 根据条件获取比较符信息分页列表
//
// @Tags　比较符信息管理
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
// @Router /flow/comparison/operators/list [get]
func (c *Controller) GetComparisonOperatorsList() rest.Result {
	var queryMdl QueryListMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if queryMdl.ClassifyId <= 0 {
		return rest.FailCustom(400, "请选择对应分类", rest.ERROR)
	}
	return c.Svc.ObtainComparisonOperatorsList(queryMdl)
}

// GetComparisonOperatorsBy
//
// @Summary 获取单个比较符信息的信息
//
// @Description 获取单个比较符信息信息
//
// @Tags　比较符信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param id path int true "比较符信息id"
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/operators/{id} [get]
func (c *Controller) GetComparisonOperatorsBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择比较符信息", rest.ERROR)
	}
	return c.Svc.ObtainComparisonOperatorsInfo(id)
}

// PostComparisonOperators
//
// @Summary 新增比较符信息
//
// @Description 新增比较符信息
//
// @Tags　比较符信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object model.ComparisonOperators
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/operators [post]
func (c *Controller) PostComparisonOperators() rest.Result {
	var processGroupInfo model.ComparisonOperators
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddComparisonOperators(processGroupInfo)
}

// PutComparisonOperators
//
// @Summary 根据id修改比较符信息信息
//
// @Description 根据id修改比较符信息信息
//
// @Tags　比较符信息管理
//
// @Accept application/json
//
// @Produce application/json
//
// @Param object model.ComparisonOperators
//
// @Success 200 {object} rest.Result
//
// @Failure 500 {object} rest.Result
//
// @Router /flow/comparison/operators [put]
func (c *Controller) PutComparisonOperators() rest.Result {
	var processGroupInfo model.ComparisonOperators
	if err := c.Ctx.ReadJSON(&processGroupInfo); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeComparisonOperators(processGroupInfo)
}

// DeleteComparisonOperators
//
// @Summary 删除比较符信息
//
// @Description 删除比较符信息
//
// @Tags　比较符信息管理
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
// @Router /flow/comparison/operators [delete]
func (c *Controller) DeleteComparisonOperators() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择比较符信息", rest.ERROR)
	}
	return c.Svc.DropComparisonOperators(delMdl.Ids)
}
