package condition_detail_groups

import (
	"reflect"

	"five.com/lk_flow/model"
	"five.com/lk_flow/utils"
	"five.com/technical_center/core_library.git/rest"
	lib "five.com/technical_center/core_library.git/utils/repo"
)

// Service 业务逻辑接口
type Service interface {
	ObtainConditionDetailGroupsList(mdl QueryListMdl) rest.Result             //根据条件获取条件详细信息分组列表
	ObtainConditionDetailGroupsInfo(id int64) rest.Result                     //获取单个条件详细信息分组的信息
	AddConditionDetailGroups(info model.ConditionDetailGroups) rest.Result    //新增条件详细信息分组
	ChangeConditionDetailGroups(info model.ConditionDetailGroups) rest.Result //根据id修改条件详细信息分组信息
	DropConditionDetailGroups(ids []int64) rest.Result                        //删除条件详细信息分组
}

type ServiceImpl struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repo: repository}
}

// ObtainConditionDetailGroupsList 根据条件获取条件详细信息分组列表
func (s ServiceImpl) ObtainConditionDetailGroupsList(mdl QueryListMdl) rest.Result {
	db := lib.ObtainCustomDb()
	items := s.repo.SelectPageConditionDetailGroupsByCondition(db, mdl)
	return utils.Result(items, "cdgDetailInfos")
}

// ObtainConditionDetailGroupsInfo 获取单个条件详细信息分组的信息
func (s ServiceImpl) ObtainConditionDetailGroupsInfo(id int64) rest.Result {
	db := lib.ObtainCustomDb()
	cdgDetailInfo := s.repo.SelectConditionDetailGroupsById(db, id)
	if reflect.DeepEqual(cdgDetailInfo, model.ConditionDetailGroups{}) {
		return rest.FailCustom(500, "查询失败", rest.ERROR)
	}
	return utils.Result(cdgDetailInfo, "cdgDetailInfo")
}

// AddConditionDetailGroups 新增条件详细信息分组
func (s ServiceImpl) AddConditionDetailGroups(info model.ConditionDetailGroups) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.InsertConditionDetailGroups(tx, info)
	return rest.SuccessCustom("新增成功", nil, rest.ERROR)
}

// ChangeConditionDetailGroups 根据id修改条件详细信息分组信息
func (s ServiceImpl) ChangeConditionDetailGroups(info model.ConditionDetailGroups) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.UpdateConditionDetailGroupsById(tx, info)
	return rest.SuccessCustom("修改成功", nil, rest.ERROR)
}

// DropConditionDetailGroups 删除条件详细信息分组
func (s ServiceImpl) DropConditionDetailGroups(ids []int64) rest.Result {
	tx := lib.ObtainCustomDbTx()
	s.repo.DeleteConditionDetailGroupsById(tx, ids)
	return rest.SuccessCustom("删除成功", nil, rest.ERROR)
}
