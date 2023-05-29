package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ComparisonOperators struct {
	req.CommonModel
	ClassifyId         int64  `json:"classifyId,omitempty"`         // 比较符分类id
	ComparisonOperator string `json:"comparisonOperator,omitempty"` // 比较符
}

// 定义列
func (c *ComparisonOperators) allColumn() []string {
	columns := []string{"classify_id", "comparison_operator"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ComparisonOperators) GetAllColumn() (result string) {
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
func (c *ComparisonOperators) TableName() string {
	return "comparison_operators"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ComparisonOperators) GetAllColumWithAlias(alias string) (result string) {
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
