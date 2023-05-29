package examine_rule

import (
	. "five.com/lk_flow/api/common"
	"five.com/lk_flow/model"
	"five.com/technical_center/core_library.git/log"
	"gorm.io/gorm"
)

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}

// SelectPageExamineRuleByCondition 根据条件获取审批规则列表
func (m MysqlRepository) SelectPageExamineRuleByCondition(db *gorm.DB, mdl QueryPageMdl) (items []model.ExamineRule, total int64) {
	db = db.Table(ExamineRule.TableName())
	if mdl.Search != "" || len(mdl.Search) > 0 {
		db = db.Where("rule_name like ?", "%"+mdl.Search+"%")
	}
	db = db.Count(&total)
	db = db.Order("created_at asc")
	mdl.PageNumber = mdl.PageNumber * mdl.PageSize
	if mdl.PageSize <= 0 {
		mdl.PageSize = 10
	}
	db.Select(ExamineRule.GetAllColumn())
	db.Limit(mdl.PageSize).Offset(mdl.PageNumber).Scan(&items)
	return
}

// SelectExamineRuleById 获取单个审批规则的信息
func (m MysqlRepository) SelectExamineRuleById(db *gorm.DB, id int64) (processGroupsInfo model.ExamineRule) {
	db.Table(ExamineRule.TableName()).Where("id = ?", id).Select(ExamineRule.GetAllColumn()).Scan(&processGroupsInfo)
	return
}

// InsertExamineRule 新增审批规则
func (m MysqlRepository) InsertExamineRule(tx *gorm.DB, info model.ExamineRule) {
	if err := tx.Table(ExamineRule.TableName()).Create(&info).Error; err != nil {
		log.Error("新增审批规则信息出错")
		panic("新增出错，数据已回滚")
	}
}

// UpdateExamineRuleById 根据id修改审批规则信息
func (m MysqlRepository) UpdateExamineRuleById(tx *gorm.DB, info model.ExamineRule) {
	if err := tx.Table(ExamineRule.TableName()).Where("id = ?", info.Id).Updates(&info).Error; err != nil {
		log.Error("修改审批规则信息出错")
		panic("修改出错，数据已回滚")
	}
}

// DeleteExamineRuleById 删除审批规则
func (m MysqlRepository) DeleteExamineRuleById(tx *gorm.DB, ids []int64) {
	if err := tx.Table(ExamineRule.TableName()).Delete(&ExamineRule, ids).Error; err != nil {
		log.Error("删除审批规则信息出错")
		panic("删除出错，数据已回滚")
	}
}
