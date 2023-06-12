package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type PanelPoint struct {
	req.CommonModel
	ProcessId     int64  `json:"processId,omitempty"`     // 流程id
	PointName     string `json:"pointName,omitempty"`     // 节点名称
	PointCode     string `json:"pointCode,omitempty"`     // 节点编号
	Scenario      int8   `json:"scenario,omitempty"`      // 节点使用场景（待定字段）
	PointType     int8   `json:"pointType,omitempty"`     // 节点类型（1 审批节点、2 抄送节点、3 子级流程、4 条件分支、5 发起人节点 6 结束节点）
	ExamineType   int8   `json:"examineType,omitempty"`   // 审批形式（1 会签 2 或签）
	Priority      int64  `json:"priority,omitempty"`      // 节点优先级
	ConditionType uint8  `json:"conditionType,omitempty"` // 条件创建类型（1 自建 2 默认条件）
}

// 定义列
func (c *PanelPoint) allColumn() []string {
	columns := []string{"process_id", "point_name", "point_code",
		"scenario", "point_type", "examine_type", "priority", "condition_type"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *PanelPoint) GetAllColumn() (result string) {
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
func (c *PanelPoint) TableName() string {
	return "panel_point"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *PanelPoint) GetAllColumWithAlias(alias string) (result string) {
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
