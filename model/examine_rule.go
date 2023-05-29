package model

import (
	"five.com/technical_center/core_library.git/rest/req"
)

type ExamineRule struct {
	req.CommonModel
	ParentRuleId int64  `json:"parentRuleId"` // 父级规则id
	RuleCode     string `json:"ruleCode"`     // 规则编号
	RuleName     string `json:"ruleName"`     // 规则名称
	RuleExplain  string `json:"ruleExplain"`  // 规则说明
	RuleType     int64  `json:"ruleType"`     // 规则类型（参与者形式规则、流程规则、节点规则、异常规则）
	RuleLevel    int64  `json:"ruleLevel"`    // 规则等级（1为最大，依次递增，权限依次递减）
}

// 定义列
func (c *ExamineRule) allColumn() []string {
	columns := []string{"parent_rule_id", "rule_code", "rule_name", "rule_explain",
		"rule_type", "rule_level"}
	commonMdl := req.CommonModel{}
	columns = append(columns, commonMdl.GetCommonModelColumns()...)
	return columns
}

// GetAllColumn 获取所有列
func (c *ExamineRule) GetAllColumn() (result string) {
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
func (c *ExamineRule) TableName() string {
	return "examine_rule"
}

// GetAllColumWithAlias 获取带有别名的所有列
func (c *ExamineRule) GetAllColumWithAlias(alias string) (result string) {
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
