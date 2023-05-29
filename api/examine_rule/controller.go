package examine_rule

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

// GetExamineRulePage
//
// @Summary 根据条件获取审批规则列表
// @Description 根据条件获取审批规则列表
// @Tags　审批规则管理
// @Accept application/json
// @Produce application/json
// @Param object QueryListMdl
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/examine/rule/page [get]
func (c *Controller) GetExamineRulePage() rest.Result {
	var queryMdl QueryPageMdl
	if err := c.Ctx.ReadQuery(&queryMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ObtainExamineRulePage(queryMdl)
}

// GetExamineRuleBy
//
// @Summary 获取单个审批规则的信息
// @Description 获取单个审批规则信息
// @Tags　审批规则管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "审批规则id"
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/examine/rule/{id} [get]
func (c *Controller) GetExamineRuleBy(id int64) rest.Result {
	if id <= 0 {
		return rest.FailCustom(400, "请选择审批规则", rest.ERROR)
	}
	return c.Svc.ObtainExamineRuleInfo(id)
}

// PostExamineRule
//
// @Summary 新增审批规则
// @Description 新增审批规则
// @Tags　审批规则管理
// @Accept application/json
// @Produce application/json
// @Param object model.ExamineRule
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/examine/rule [get]
func (c *Controller) PostExamineRule() rest.Result {
	var info model.ExamineRule
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.AddExamineRule(info)
}

// PutExamineRule
//
// @Summary 根据id修改审批规则信息
// @Description 根据id修改审批规则信息
// @Tags　审批规则管理
// @Accept application/json
// @Produce application/json
// @Param object model.ExamineRule
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/examine/rule [get]
func (c *Controller) PutExamineRule() rest.Result {
	var info model.ExamineRule
	if err := c.Ctx.ReadJSON(&info); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	return c.Svc.ChangeExamineRule(info)
}

// DeleteExamineRule
//
// @Summary 删除审批规则
// @Description 删除审批规则
// @Tags　审批规则管理
// @Accept application/json
// @Produce application/json
// @Param object req.DelModel
// @Success 200 {object} rest.Result
// @Failure 500 {object} rest.Result
// @Router /flow/examine/rule [get]
func (c *Controller) DeleteExamineRule() rest.Result {
	var delMdl req.DelModel
	if err := c.Ctx.ReadJSON(&delMdl); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if len(delMdl.Ids) <= 0 {
		return rest.FailCustom(400, "请选择审批规则", rest.ERROR)
	}
	return c.Svc.DropExamineRule(delMdl.Ids)
}
