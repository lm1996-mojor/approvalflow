package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type PointValue struct {
	req.CommonModel
	ApprovalCode string `json:"approvalCode,omitempty"` // 审批编号(32位)
	PointId      int64  `json:"pointId,omitempty"`      // 节点id
	PointRate    uint8  `json:"pointRate,omitempty"`    // 节点进度（通过、退回、驳回）
	NextStep     int64  `json:"nextStep,omitempty"`     // 下一节点
	NextStepType int8   `json:"nextStepType,omitempty"` // 下一节点类型1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点
}

// 定义列
func (c *PointValue) allColumn() []string {
	columns := []string{"approval_code", "point_id", "point_rate", "next_step", "next_step_type"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *PointValue) GetAllColumn() (result string) {
	columns := c.allColumn()
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			result += columns[i]
		} else {
			result += columns[i] + " , "
		}
	}
	return result
}

// TableName 获取表名
func (c *PointValue) TableName() string {
	return "point_value"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *PointValue) GetAllColumWithAlias(alias string) (result string) {
	columns := c.allColumn()
	for i := 0; i < len(columns); i++ {
		if i == len(columns)-1 {
			result += alias + "." + columns[i]
		} else {
			result += alias + "." + columns[i] + " , "
		}
	}
	return result
}
