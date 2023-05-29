package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

// Participant 参与者信息结构体
type Participant struct {
	req.CommonModel
	ObjId          int64  `json:"objId,omitempty"`          // 参与者id
	PointValueId   int64  `json:"pointValueId,omitempty"`   // 节点值id
	OrderNo        int64  `json:"orderNo,omitempty"`        // 参与者顺序
	ApprovalResult int8   `json:"approvalResult,omitempty"` // 审批结果
	Opinions       string `json:"opinions,omitempty"`       // 审批意见
}

// 定义列
func (c *Participant) allColumn() []string {
	columns := []string{"obj_id", "point_value_id", "order_no", "approval_result", "opinions"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *Participant) GetAllColumn() (result string) {
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
func (c *Participant) TableName() string {
	return "participant"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *Participant) GetAllColumWithAlias(alias string) (result string) {
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
