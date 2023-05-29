package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ProcessGroups struct {
	req.CommonModel
	AppCode      string `json:"appCode,omitempty"`      // 应用编码
	BusinessCode string `json:"businessCode,omitempty"` // 业务编码
	ClientId     int64  `json:"-,omitempty"`            // 租户id
	GroupName    string `json:"groupName,omitempty"`    // 分组名称
	OrderNo      int64  `json:"orderNo,omitempty"`      // 分组排序
	IsDefault    uint8  `json:"isDefault,omitempty"`    // 是否默认（1 是 2 否）
}

// 定义列
func (c *ProcessGroups) allColumn() []string {
	columns := []string{"app_code", "business_code", "client_id", "group_name", "order_no", "is_default"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ProcessGroups) GetAllColumn() (result string) {
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
func (c *ProcessGroups) TableName() string {
	return "process_groups"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ProcessGroups) GetAllColumWithAlias(alias string) (result string) {
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
