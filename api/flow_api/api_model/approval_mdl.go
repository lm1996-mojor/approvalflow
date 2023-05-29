package api_model

import (
	"five.com/lk_flow/model"
)

type ApprovalParams struct {
	model.ProcessValue
	CtlValues    []model.CtlValueInfo `json:"ctlValues"`    // 表单值信息
	PointDetails []PointDetail        `json:"pointDetails"` // 节点信息
}

type PointDetail struct {
	model.PointValue
	PointType        int8                `json:"pointType"`        // 节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）
	ExamineType      int8                `json:"examineType"`      // 审批形式（1 依次审批 2 会签 3 或签）
	ParticipantInfos []model.Participant `json:"participantInfos"` // 参与者信息
}

// CurrentApprovalStatus 当前审批结果
type CurrentApprovalStatus struct {
	AppCode      string            `json:"appCode"`      // 应用编码
	BusinessCode string            `json:"businessCode"` // 业务编码
	ClientId     int64             `json:"-"`            // 租户id
	ApprovalCode string            `json:"approvalCode"` // 审批编号(32位)
	NodeDetailId int64             `json:"nodeDetailId"` // 节点值id（节点值）
	NodeType     int8              `json:"nodeType"`     // 节点类型
	ExamineType  int8              `json:"examineType"`  // 审批形式（1 会签 2 或签）
	Participant  model.Participant `json:"participants"` // 审批参与人信息
	BackNodeId   int64             `json:"backNodeId"`   // 退回到的节点
}

type QueryApprovalParam struct {
	AppCode      string `json:"appCode"`      // 应用编码
	BusinessCode string `json:"businessCode"` // 业务编码
	ClientId     int64  `json:"-"`            // 租户id
	ApprovalCode string `json:"approvalCode"` // 审批编号(32位)
	ProcessId    int64  `json:"processId"`    // 流程id
	UserId       int64  `json:"userId"`       // 用户id(即参与人id)
	QueryType    int8   `json:"queryType"`    // 查询类型（1 待处理 2 已处理 3 已发起 4 收到的）
}
