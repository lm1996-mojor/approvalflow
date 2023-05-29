package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ParticipantFormat struct {
	req.CommonModel
	FormatName    string `json:"formatName"`    // 形式名称
	FormatExplain string `json:"formatExplain"` // 形式说明
}

// 定义列
func (c *ParticipantFormat) allColumn() []string {
	columns := []string{"format_name", "format_explain"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ParticipantFormat) GetAllColumn() (result string) {
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
func (c *ParticipantFormat) TableName() string {
	return "participant_format"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ParticipantFormat) GetAllColumWithAlias(alias string) (result string) {
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
