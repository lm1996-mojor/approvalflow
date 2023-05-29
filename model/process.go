package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type Process struct {
	req.CommonModel
	GroupId          int64  `json:"groupId,omitempty"`          // 流程分组id
	FlowName         string `json:"flowName,omitempty"`         // 流程名
	FlowCode         string `json:"flowCode,omitempty"`         // 流程编号
	Icon             string `json:"icon,omitempty"`             // 流程图标
	Illustrate       string `json:"illustrate,omitempty"`       // 流程说明
	SubsidiaryFormId int64  `json:"subsidiaryFormId,omitempty"` // 流程附属表单（表单来源为外部来源时才有值）
	FormSource       uint8  `json:"formSource,omitempty"`       // 流程表单来源（流程自建、外部来源）
	ProcessStatus    int8   `json:"processStatus,omitempty"`    // 流程状态(1 已发布 2 未发布)
}

// 定义列
func (c *Process) allColumn() []string {
	columns := []string{"group_id", "flow_name", "flow_code",
		"icon", "illustrate", "subsidiary_form_id", "form_source", "process_status"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *Process) GetAllColumn() (result string) {
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
func (c *Process) TableName() string {
	return "process"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *Process) GetAllColumWithAlias(alias string) (result string) {
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
