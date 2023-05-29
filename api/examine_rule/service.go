package examine_rule

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainExamineRulePage(mdl QueryPageMdl) rest.Result   //根据条件获取审批规则列表
	ObtainExamineRuleInfo(id int64) rest.Result           //获取单个审批规则的信息
	AddExamineRule(info model.ExamineRule) rest.Result    //新增审批规则
	ChangeExamineRule(info model.ExamineRule) rest.Result //根据id修改审批规则信息
	DropExamineRule(ids []int64) rest.Result              //删除审批规则
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainExamineRulePage 根据条件获取审批规则列表
func (s ServiceImpl) ObtainExamineRulePage(mdl QueryPageMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items, total := s.repo.SelectPageExamineRuleByCondition(db, mdl)
	return rest.NewQueryPage(items, mdl.PageNumber, mdl.PageSize, total)
}

// ObtainExamineRuleInfo 获取单个审批规则的信息
func (s ServiceImpl) ObtainExamineRuleInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	examineRuleInfo := s.repo.SelectExamineRuleById(db, id)
	if reflect.DeepEqual(examineRuleInfo, model.ExamineRule{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(examineRuleInfo, "examineRuleInfo")
}

// AddExamineRule 新增审批规则
func (s ServiceImpl) AddExamineRule(info model.ExamineRule) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertExamineRule(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeExamineRule 根据id修改审批规则信息
func (s ServiceImpl) ChangeExamineRule(info model.ExamineRule) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateExamineRuleById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropExamineRule 删除审批规则
func (s ServiceImpl) DropExamineRule(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteExamineRuleById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}
