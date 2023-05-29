package rel

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ProcessObjRuleRel struct {
	Id           int64 `json:"id"`           // 主键id
	RuleId       int64 `json:"ruleId"`       // 规则id
	ProcessObjId int64 `json:"processObjId"` // 路程对象id
	OwnerType    int64 `json:"ownerType"`    // 关系所属对象类型（1 参与者、2 流程主体、3 节点）
}

// 定义列
func (c *ProcessObjRuleRel) allColumn() []string {
	columns := []string{"id", "rule_id", "process_obj_id", "owner_type"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ProcessObjRuleRel) GetAllColumn() (result string) {
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
func (c *ProcessObjRuleRel) TableName() string {
	return "process_obj_rule_rel"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ProcessObjRuleRel) GetAllColumWithAlias(alias string) (result string) {
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
