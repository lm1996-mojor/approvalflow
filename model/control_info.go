package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ControlInfo struct {
	req.CommonModel
	ParentId      int64  `json:"parentId,omitempty"`      // 父级id（关联子集表和规则表id）
	TabId         int64  `json:"tabId,omitempty"`         // 标签id
	OwnerType     uint8  `json:"ownerType,omitempty"`     // 控件所属类型（1 流程主体 2 流程规则）
	CnName        string `json:"cnName,omitempty"`        // 控件中文名（唯一）
	EnName        string `json:"enName,omitempty"`        // 控件英文名（唯一）
	CtlCode       string `json:"ctlCode,omitempty"`       // 控件编码（唯一）
	Enable        uint8  `json:"enable,omitempty"`        // 是否开启（1 开启 2 禁用）
	Required      uint8  `json:"required,omitempty"`      // 控件值是否必填（1 是 2 否）
	FieldName     string `json:"fieldName,omitempty"`     // 控件数据库表列名（唯一）
	ComponentType string `json:"componentType,omitempty"` // 控件类型
	ValueType     string `json:"valueType,omitempty"`     // 控件值类型
	Props         string `json:"props,omitempty"`         // 控件属性
	OrderNo       int64  `json:"orderNo,omitempty"`       // 控件排序
	IsDefault     uint8  `json:"isDefault,omitempty"`     // 是否为默认控件（1 是 2 否）
}

// 定义列
func (c *ControlInfo) allColumn() []string {

	columns := []string{"parent_id", "tab_id", "owner_type", "cn_name", "en_name", "ctl_code", "enable",
		"required", "field_name", "component_type", "value_type", "props", "order_no", "is_default"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ControlInfo) GetAllColumn() (result string) {
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
func (c *ControlInfo) TableName() string {
	return "control_info"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ControlInfo) GetAllColumWithAlias(alias string) (result string) {
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
