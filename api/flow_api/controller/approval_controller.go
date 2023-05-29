package controller

import (
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/api/flow_api/service"
	"five.com/technical_center/core_library.git/rest"
	"github.com/kataras/iris/v12"
)

type ApprovalController struct {
	Ctx iris.Context
	Svc service.ApprovalServiceService
}

func ApprovalNewController(repository repository.ProcessValueRepository) *ApprovalController {
	return &ApprovalController{Svc: service.NewService(repository)}
}

// GetUserApprovalInfo
//
// @Summary 获取用户的审批信息
//
// @Description 获取用户的审批信息
//
// @Group　审批开放api
//
// @Accept application/json
//
// @Produce application/json
//
// @Param
//
// @Success
//
// @Failure
//
// @Router /flow/open/api/user/approval/info
func (c *ApprovalController) GetUserApprovalInfo() rest.Result {
	var params api_model.QueryApprovalParam
	if err := c.Ctx.ReadQuery(&params); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if params.ApprovalCode == "" || len(params.ApprovalCode) <= 0 {
		return rest.FailCustom(400, "请指定审批流程", rest.ERROR)
	}
	if params.AppCode == "" || len(params.AppCode) <= 0 {
		return rest.FailCustom(400, "请确认应用", rest.ERROR)
	}
	return c.Svc.FindUserApprovalInfo(params)
}

// PostInitiateApprovalFlow
//
// @Summary 发起审批流程
//
// @Description 发起审批流程
//
// @Group　审批开放api
//
// @Accept application/json
//
// @Produce application/json
//
// @Param
//
// @Success
//
// @Failure
//
// @Router /flow/open/api/initiate/approval/flow
func (c *ApprovalController) PostInitiateApprovalFlow() rest.Result {
	var params api_model.ApprovalParams
	if err := c.Ctx.ReadJSON(&params); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if params.ApprovalCode == "" || len(params.ApprovalCode) <= 0 {
		return rest.FailCustom(400, "请指定审批流程", rest.ERROR)
	}
	if params.AppCode == "" || len(params.AppCode) <= 0 {
		return rest.FailCustom(400, "请确认应用", rest.ERROR)
	}
	return c.Svc.InitiateApproval(params)
}

// PutApproval
//
// @Summary 参与人审批
//
// @Description 参与人审批
//
// @Group 审批开放api
//
// @Accept application/json
//
// @Produce application/json
//
// @Param
//
// @Success
//
// @Failure
//
// @Router /flow/open/api/approval
func (c *ApprovalController) PutApproval() rest.Result {
	var params api_model.CurrentApprovalStatus
	if err := c.Ctx.ReadJSON(&params); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if params.ApprovalCode == "" || len(params.ApprovalCode) <= 0 {
		return rest.FailCustom(400, "请指定审批流程", rest.ERROR)
	}
	if params.AppCode == "" || len(params.AppCode) <= 0 {
		return rest.FailCustom(400, "请确认应用", rest.ERROR)
	}
	if params.Participant.ApprovalResult > 3 {
		return rest.FailCustom(400, "请进行有效审批", rest.ERROR)
	}
	if params.Participant.ApprovalResult <= 3 && params.Participant.ApprovalResult > 1 && len(params.Participant.Opinions) <= 0 {
		return rest.FailCustom(400, "驳回或者退回时必须填写审批意见", rest.ERROR)
	}
	return c.Svc.Approval(params)
}

// DeleteApproval
//
// @Summary 撤回审批流程
//
// @Description 撤回审批流程
//
// @Group 审批开放api
//
// @Accept application/json
//
// @Produce application/json
//
// @Param
//
// @Success
//
// @Failure
//
// @Router /flow/open/api/approval
func (c *ApprovalController) DeleteApproval() rest.Result {
	var params api_model.CurrentApprovalStatus
	if err := c.Ctx.ReadJSON(&params); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if params.ApprovalCode == "" || len(params.ApprovalCode) <= 0 {
		return rest.FailCustom(400, "请指定审批流程", rest.ERROR)
	}
	if params.AppCode == "" || len(params.AppCode) <= 0 {
		return rest.FailCustom(400, "请确认应用", rest.ERROR)
	}
	return c.Svc.RevokeApprovalProcess(params)
}
