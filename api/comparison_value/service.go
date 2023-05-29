package comparison_value

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainComparisonValueList(mdl QueryListMdl) rest.Result       //根据条件获取比较值列表
	ObtainComparisonValueInfo(id int64) rest.Result               //获取单个比较值的信息
	AddComparisonValue(info model.ComparisonValue) rest.Result    //新增比较值
	ChangeComparisonValue(info model.ComparisonValue) rest.Result //根据id修改比较值信息
	DropComparisonValue(ids []int64) rest.Result                  //删除比较值
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainComparisonValueList 根据条件获取比较值列表
func (s ServiceImpl) ObtainComparisonValueList(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items := s.repo.SelectPageComparisonValueByCondition(db, mdl)
	return utils.Result(items, "cvInfos")
}

// ObtainComparisonValueInfo 获取单个比较值的信息
func (s ServiceImpl) ObtainComparisonValueInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	cdgInfo := s.repo.SelectComparisonValueById(db, id)
	if reflect.DeepEqual(cdgInfo, model.ComparisonValue{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(cdgInfo, "cvInfo")
}

// AddComparisonValue 新增比较值
func (s ServiceImpl) AddComparisonValue(info model.ComparisonValue) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertComparisonValue(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeComparisonValue 根据id修改比较值信息
func (s ServiceImpl) ChangeComparisonValue(info model.ComparisonValue) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateComparisonValueById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropComparisonValue 删除比较值
func (s ServiceImpl) DropComparisonValue(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteComparisonValueById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}
