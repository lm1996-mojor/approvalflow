package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ComparisonClassify struct {
	req.CommonModel
	ClassifyName string `json:"classifyName,omitempty"` //分类名
	ClassifyChar string `json:"classifyChar,omitempty"` //分类标识（int、float、str）
}

// 定义列
func (c *ComparisonClassify) allColumn() []string {
	columns := []string{"classify_name", "classify_char"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ComparisonClassify) GetAllColumn() (result string) {
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
func (c *ComparisonClassify) TableName() string {
	return "comparison_classify"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ComparisonClassify) GetAllColumWithAlias(alias string) (result string) {
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
