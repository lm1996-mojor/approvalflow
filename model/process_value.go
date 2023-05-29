package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ProcessValue struct {
	req.CommonModel
	AppCode      string `json:"appCode,omitempty"`      // 应用编码
	BusinessCode string `json:"businessCode,omitempty"` // 业务编码
	ClientId     int64  `json:"-,omitempty"`            // 租户id
	ProcessId    int64  `json:"processId,omitempty"`    // 流程id
	ApprovalCode string `json:"approvalCode,omitempty"` // 审批编号(32位)
	ProcessRate  uint8  `json:"processRate,omitempty"`  // 流程结果进度（未开始、进行中、通过、退回、驳回）
}

// 定义列
func (c *ProcessValue) allColumn() []string {
	columns := []string{"app_code", "business_code", "client_id", "process_id", "approval_code", "process_rate"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ProcessValue) GetAllColumn() (result string) {
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
func (c *ProcessValue) TableName() string {
	return "process_value"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ProcessValue) GetAllColumWithAlias(alias string) (result string) {
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
