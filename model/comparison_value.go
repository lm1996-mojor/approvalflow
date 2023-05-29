package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ComparisonValue struct {
	req.CommonModel
	CondDetailInfoId int64  `json:"condDetailInfoId,omitempty"` // 条件详细信息id
	CompValue        string `json:"compValue,omitempty"`        // 比较值
}

// 定义列
func (c *ComparisonValue) allColumn() []string {
	columns := []string{"cond_detail_info_id", "comp_value"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ComparisonValue) GetAllColumn() (result string) {
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
func (c *ComparisonValue) TableName() string {
	return "comparison_value"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ComparisonValue) GetAllColumWithAlias(alias string) (result string) {
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
