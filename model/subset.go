package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type Subset struct {
	req.CommonModel
	ProcessId int64  `json:"processId,omitempty"` // 流程id
	ParentId  int64  `json:"parentId,omitempty"`  // 父级子集id
	CnName    string `json:"cnName,omitempty"`    // 子集中文名（唯一）
	EnName    string `json:"enName,omitempty"`    // 子集英文名（唯一）
	SubCode   string `json:"subCode,omitempty"`   // 子集编码（唯一）
	Enable    int8   `json:"enable,omitempty"`    // 是否开启（1 开启 2 禁用）
	FieldName string `json:"fieldName,omitempty"` // 子集数据库表列名（唯一）
	Props     string `json:"props,omitempty"`     // 子集属性
	OrderNo   int32  `json:"orderNo,omitempty"`   // 子集排序
	IsDefault int8   `json:"isDefault,omitempty"` // 是否为默认子集（1 是 2 否）
}

// 定义列
func (c *Subset) allColumn() []string {
	columns := []string{"process_id", "parent_id", "cn_name", "en_name", "sub_code", "enable",
		"field_name", "props", "order_no", "is_default"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *Subset) GetAllColumn() (result string) {
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
func (c *Subset) TableName() string {
	return "subset"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *Subset) GetAllColumWithAlias(alias string) (result string) {
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
