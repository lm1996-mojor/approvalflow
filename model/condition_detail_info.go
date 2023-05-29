package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ConditionDetailInfo struct {
	req.CommonModel
	GroupsId int64 `json:"groupsId,omitempty"` // 分组id
	CtlId    int64 `json:"ctlId,omitempty"`    // 控件id
}

// 定义列
func (c *ConditionDetailInfo) allColumn() []string {
	columns := []string{"groups_id", "ctl_id"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ConditionDetailInfo) GetAllColumn() (result string) {
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
func (c *ConditionDetailInfo) TableName() string {
	return "condition_detail_info"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ConditionDetailInfo) GetAllColumWithAlias(alias string) (result string) {
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
