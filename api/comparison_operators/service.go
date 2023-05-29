package comparison_operators

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainComparisonOperatorsList(mdl QueryListMdl) rest.Result           //根据条件获取比较符列表
	ObtainComparisonOperatorsInfo(id int64) rest.Result                   //获取单个比较符的信息
	AddComparisonOperators(info model.ComparisonOperators) rest.Result    //新增比较符
	ChangeComparisonOperators(info model.ComparisonOperators) rest.Result //根据id修改比较符信息
	DropComparisonOperators(ids []int64) rest.Result                      //删除比较符
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainComparisonOperatorsList 根据条件获取比较符列表
func (s ServiceImpl) ObtainComparisonOperatorsList(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items := s.repo.SelectPageComparisonOperatorsByCondition(db, mdl)
	return utils.Result(items, "coInfos")
}

// ObtainComparisonOperatorsInfo 获取单个比较符的信息
func (s ServiceImpl) ObtainComparisonOperatorsInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	cdgInfo := s.repo.SelectComparisonOperatorsById(db, id)
	if reflect.DeepEqual(cdgInfo, model.ComparisonOperators{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(cdgInfo, "coInfo")
}

// AddComparisonOperators 新增比较符
func (s ServiceImpl) AddComparisonOperators(info model.ComparisonOperators) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertComparisonOperators(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeComparisonOperators 根据id修改比较符信息
func (s ServiceImpl) ChangeComparisonOperators(info model.ComparisonOperators) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateComparisonOperatorsById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropComparisonOperators 删除比较符
func (s ServiceImpl) DropComparisonOperators(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteComparisonOperatorsById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}
