package controller

import (
	"five.com/lk_flow/api/flow_api/api_model"
	"five.com/lk_flow/api/flow_api/repository"
	"five.com/lk_flow/api/flow_api/service"
	"five.com/technical_center/core_library.git/log"
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

// GetApprovalPage
//
// @Summary 获取审批流程分页列表
//
// @Description 获取审批流程分页列表
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
// @Router /flow/open/api/approval/page
func (c *ApprovalController) GetApprovalPage() rest.Result {
	var params api_model.QueryApprovalParam
	if err := c.Ctx.ReadQuery(&params); err != nil {
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if params.AppCode == "" || len(params.AppCode) <= 0 {
		return rest.FailCustom(400, "请确认应用", rest.ERROR)
	}
	if params.PageNumber < 0 {
		params.PageNumber = 0
	}
	if params.PageNumber < 10 {
		params.PageSize = 10
	}
	return c.Svc.ObtainUserApprovalInfoPage(params)
}

// GetApprovalInfoBy
//
// @Summary 根据审批编号获取单个审批信息
//
// @Description 根据审批编号获取单个审批信息
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
// @Router /flow/open/api/approval/info/{approvalCode}
func (c *ApprovalController) GetApprovalInfoBy(approvalCode string) rest.Result {
	if approvalCode == "" {
		return rest.FailCustom(400, "请确认审批流程", rest.ERROR)
	}
	return c.Svc.ObtainSingleApprovalInfoByApprovalCode(approvalCode)
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
		log.Error(err.Error())
		return rest.FailCustom(500, "参数解析错误", rest.ERROR)
	}
	if params.AppCode == "" || len(params.AppCode) <= 0 {
		return rest.FailCustom(400, "请确认应用", rest.ERROR)
	}

	if len(params.PointDetails) <= 3 {
		return rest.FailCustom(400, "节点存在问题，少于三个节点", rest.ERROR)
	}
	promoterFlag := false
	endFlag := false
	otherFlag := false
	for _, detail := range params.PointDetails {
		if detail.PointType < 5 {
			otherFlag = true
		}
		if detail.PointType == 5 {
			promoterFlag = true
		}
		if detail.PointType == 6 {
			endFlag = true
		}
	}
	if !(promoterFlag && endFlag && otherFlag) {
		return rest.FailCustom(400, "节点不规范，最少要拥有三个节点，（发起人、结束节点、审批/抄送/子流程/条件等节点）", rest.ERROR)
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
	if params.Participant.ApprovalResult > 3 && params.Participant.ApprovalResult > 0 {
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
