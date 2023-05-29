package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type CtlTab struct {
	req.CommonModel
	TabName string `json:"tabName,omitempty"` // 标签名称
}

// 定义列
func (c *CtlTab) allColumn() []string {
	columns := []string{"tab_name"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *CtlTab) GetAllColumn() (result string) {
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
func (c *CtlTab) TableName() string {
	return "ctl_tab"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *CtlTab) GetAllColumWithAlias(alias string) (result string) {
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
