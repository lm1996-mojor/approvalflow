package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ConditionDetailGroups struct {
	req.CommonModel
	PointId   int64  `json:"pointId,omitempty"`   // 条件主体信息id
	GroupName string `json:"groupName,omitempty"` // 分组名称
	GroupCode string `json:"groupCode,omitempty"` // 分组编号
	PassAsk   uint8  `json:"passAsk,omitempty"`   // 通过要求（1 满足全部 2 其中之一）
}

// 定义列
func (c *ConditionDetailGroups) allColumn() []string {
	columns := []string{"point_id", "group_name", "group_code", "pass_ask"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ConditionDetailGroups) GetAllColumn() (result string) {
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
func (c *ConditionDetailGroups) TableName() string {
	return "condition_detail_groups"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ConditionDetailGroups) GetAllColumWithAlias(alias string) (result string) {
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
