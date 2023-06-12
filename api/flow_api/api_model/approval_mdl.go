package api_model

import (
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/rest/req"
)

// ApprovalParams 审批流程整体数据存储对象
type ApprovalParams struct {
	model.ProcessValue
	CtlValues    []CtlDetail   `json:"ctlValues,omitempty"`    // 表单值信息
	PointDetails []PointDetail `json:"pointDetails,omitempty"` // 节点信息
}

type CtlDetail struct {
	model.CtlValueInfo
	ParentId      int64  `json:"parentId,omitempty"`      // 父级id（关联子集表和规则表id）
	TabId         int64  `json:"tabId,omitempty"`         // 标签id
	OwnerType     uint8  `json:"ownerType,omitempty"`     // 控件所属类型（1 流程主体 2 流程规则）
	CnName        string `json:"cnName,omitempty"`        // 控件中文名（唯一）
	EnName        string `json:"enName,omitempty"`        // 控件英文名（唯一）
	CtlCode       string `json:"ctlCode,omitempty"`       // 控件编码（唯一）
	Enable        uint8  `json:"enable,omitempty"`        // 是否开启（1 开启 2 禁用）
	Required      uint8  `json:"required,omitempty"`      // 控件值是否必填（1 是 2 否）
	FieldName     string `json:"fieldName,omitempty"`     // 控件数据库表列名（唯一）
	ComponentType string `json:"componentType,omitempty"` // 控件类型
	ValueType     string `json:"valueType,omitempty"`     // 控件值类型
	Props         string `json:"props,omitempty"`         // 控件属性
	OrderNo       int64  `json:"orderNo,omitempty"`       // 控件排序
	IsDefault     uint8  `json:"isDefault,omitempty"`     // 是否为默认控件（1 是 2 否）
}

type PointDetail struct {
	PointValueDetail
	ParticipantInfos []model.Participant `json:"participantInfos"` // 参与者信息
}

type PointValueDetail struct {
	model.PointValue
	PointType   int8   `json:"pointType"`   // 节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）
	ExamineType int8   `json:"examineType"` // 审批形式（1 会签 2 或签）
	PointName   string `json:"pointName"`   // 节点名称
}

// CurrentApprovalStatus 当前审批结果
type CurrentApprovalStatus struct {
	AppCode             string            `json:"appCode"`             // 应用编码
	ClientId            int64             `json:"-"`                   // 租户id
	ApprovalCode        string            `json:"approvalCode"`        // 审批编号(32位)
	NodeDetailId        int64             `json:"nodeDetailId"`        // 节点值id（节点值）
	NodeType            int8              `json:"nodeType"`            // 节点类型
	ExamineType         int8              `json:"examineType"`         // 审批形式（1 会签 2 或签）
	Participant         model.Participant `json:"participants"`        // 审批参与人信息
	BackNodeId          int64             `json:"backNodeId"`          // 退回到的节点
	BackIntervalNodeIds []int64           `json:"backIntervalNodeIds"` // 退回区间的id数组
}

type QueryApprovalParam struct {
	req.PageParam
	AppCode      string `json:"appCode"`      // 应用编码 -- 必填
	ClientId     int64  `json:"-"`            // 租户id
	ApprovalCode string `json:"approvalCode"` // 审批编号(32位)
	ProcessId    int64  `json:"processId"`    // 流程id
	UserId       int64  `json:"userId"`       // 用户id(即参与人id)
	QueryType    int    `json:"queryType"`    // 查询类型（1 待处理 2 已处理 3 已发起 4 收到的）
}

// 响应

// RepoApprovalParam 审批流程响应结构体
type RepoApprovalParam struct {
	ApprovalId            int64               `json:"approvalId"`            // 审批id
	ApprovalCode          string              `json:"approvalCode"`          // 审批编号(32位)
	ProcessId             int64               `json:"processId"`             // 流程id
	NodeDetailId          int64               `json:"nodeDetailId"`          // 当前审批节点值id（节点值）
	ApprovalTitle         string              `json:"approvalTitle"`         // 审批标题
	ExamineType           int8                `json:"examineType"`           // 审批形式（1 会签 2 或签）
	PointType             int8                `json:"pointType"`             // 节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）
	PointRate             int8                `json:"pointRate"`             // 节点进度（1 同意 2退回 3驳回 4审批中 5待执行 6撤销）
	PointName             string              `json:"pointName"`             // 节点名称
	CurrentApprovalResult string              `json:"currentApprovalResult"` // 当前审批结果描述
	CtlValues             []CtlDetail         `json:"ctlValues"`             // 表单值信息
	Participant           []model.Participant `json:"participants"`          // 审批参与人信息
	ApprovalIsEnd         bool                `json:"approvalIsEnd"`         // 审批是否结束
}

// RepoSingleApprovalParam 单个审批流程响应结构体
type RepoSingleApprovalParam struct {
	ApprovalInfo        ApprovalParams    `json:"approvalInfo"`        // 审批流程整体信息
	CurrentApprovalInfo RepoApprovalParam `json:"currentApprovalInfo"` // 该审批流程当前的进度
}
