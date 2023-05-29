package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type CtlValueInfo struct {
	req.CommonModel
	ApprovalCode string `json:"approvalCode,omitempty"` // 审批编号
	CtlId        int64  `json:"ctlId,omitempty"`        // 控件id
	CtlValue     string `json:"ctlValue,omitempty"`     // 控件值信息
}

// 定义列
func (c *CtlValueInfo) allColumn() []string {
	columns := []string{"approval_code", "ctl_id", "ctl_value"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *CtlValueInfo) GetAllColumn() (result string) {
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
func (c *CtlValueInfo) TableName() string {
	return "ctl_value_info"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *CtlValueInfo) GetAllColumWithAlias(alias string) (result string) {
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
