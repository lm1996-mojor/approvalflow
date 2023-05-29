package examine_rule

import (
	"five.com/lk_flow/model"
	"gorm.io/gorm"
)

// Repository 操作数据库接口
type Repository interface {
	SelectPageExamineRuleByCondition(db *gorm.DB, mdl QueryPageMdl) ([]model.ExamineRule, int64) //根据条件获取审批规则列表
	SelectExamineRuleById(db *gorm.DB, id int64) model.ExamineRule                               //获取单个审批规则的信息
	InsertExamineRule(tx *gorm.DB, info model.ExamineRule)                                       //新增审批规则
	UpdateExamineRuleById(tx *gorm.DB, info model.ExamineRule)                                   //根据id修改审批规则信息
	DeleteExamineRuleById(tx *gorm.DB, ids []int64)                                              //删除审批规则
}
